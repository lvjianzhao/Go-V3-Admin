package system

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"

	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
)

type ApiService struct{}

// AddApi 添加api
func (a *ApiService) AddApi(api systemModel.ApiModel) (*systemModel.ApiModel, error) {
	var err error
	tx := global.DB.Begin()
	if api.Method != "" {
		methods := strings.Split(api.Method, ",")
		for _, method := range methods {
			// 在这里处理每个 method 的逻辑
			if !errors.Is(tx.Where("path = ? AND method = ?", api.Path, method).First(&systemModel.ApiModel{}).Error, gorm.ErrRecordNotFound) {
				//return nil, errors.New("存在相同api")
				msg := fmt.Sprintf("path: %s,method: %s,warn: 存在相同的API，跳过...", api.Path, api.Method)
				global.LOG.Warn(msg)
				break
			}
			api.Method = method
			err = tx.Create(&api).Error
			api.ID += 1
			if err != nil {
				msg := fmt.Sprintf("path: %s,method: %s,API添加失败,error: %s", api.Path, api.Method, err)
				global.LOG.Warn(msg)
				tx.Rollback()
				return &api, err
			}
		}
	}
	// 自动给admin角色增加权限
	var roleModel systemModel.RoleModel
	err = tx.Model(&roleModel).Where("role_name = ?", "admin").First(&roleModel).Error
	if err != nil {
		global.LOG.Error("查询admin的id失败,", zap.Error(err))
	}
	var allApiModel []systemReq.CasbinInfo
	err = tx.Model(&systemModel.ApiModel{}).Select("path", "method").Find(&allApiModel).Error
	if err != nil {
		global.LOG.Error("AddApi 查询所有api", zap.Error(err))
		tx.Rollback()
		return &api, err
	}

	var cs CasbinService
	err = cs.EditCasbin(roleModel.ID, allApiModel)
	if err != nil {
		global.LOG.Error("新增api后自动授权给admin失败", zap.Error(err))
		tx.Rollback()
		return &api, err
	}
	tx.Commit()
	return &api, err
}

// GetApis 获取所有api
func (a *ApiService) GetApis(apiSp systemReq.ApiSearchParams) ([]systemModel.ApiModel, int64, error) {
	limit := apiSp.PageSize
	offset := apiSp.PageSize * (apiSp.Page - 1)
	db := global.DB.Model(&systemModel.ApiModel{})
	var apiList []systemModel.ApiModel

	if apiSp.Path != "" {
		db = db.Where("path LIKE ?", "%"+apiSp.Path+"%")
	}

	if apiSp.Description != "" {
		db = db.Where("description LIKE ?", "%"+apiSp.Description+"%")
	}

	if apiSp.Method != "" {
		apiMethodSlice := strings.Split(apiSp.Method, ",")
		orConditions := make([]string, len(apiMethodSlice))
		args := make([]interface{}, len(apiMethodSlice))
		for i, method := range apiMethodSlice {
			// 在这里处理每个 method 的逻辑
			orConditions[i] = "method = ?"
			args[i] = method
		}
		db = db.Where(strings.Join(orConditions, " OR "), args...)
	}

	if apiSp.ApiGroup != "" {
		apiGroupSlice := strings.Split(apiSp.ApiGroup, ",")
		orConditions := make([]string, len(apiGroupSlice))
		args := make([]interface{}, len(apiGroupSlice))
		for i, group := range apiGroupSlice {
			// 在这里处理每个 api group 的逻辑
			orConditions[i] = "api_group = ?"
			args[i] = group
		}
		db = db.Where(strings.Join(orConditions, " OR "), args...)
	}

	var total int64
	err := db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if apiSp.OrderKey != "" {
			var orderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			orderMap["updated_at"] = true
			if orderMap[apiSp.OrderKey] {
				if apiSp.Desc {
					orderStr = apiSp.OrderKey + " desc"
				} else {
					orderStr = apiSp.OrderKey
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", apiSp.OrderKey)
				return apiList, total, err
			}

			err = db.Order(orderStr).Find(&apiList).Error
		} else {
			err = db.Find(&apiList).Error
		}
	}
	return apiList, total, err
}

// GetApiGroups 获取所有API分组
func (a *ApiService) GetApiGroups() (apiGroups []string, total int64, error error) {
	db := global.DB.Model(&systemModel.ApiModel{})
	// 查询去重的 ApiGroup 列
	err := db.Distinct("api_group").Pluck("api_group", &apiGroups).Error
	if err != nil {
		return nil, 0, err
	}
	total = int64(len(apiGroups))
	return
}

// GetElTreeApis 获取所有api tree
// element-plus el-tree的数据格式
func (a *ApiService) GetElTreeApis(roleId uint) (list []systemModel.ApiTree, checkedKey []string, err error) {
	var apiModels []systemModel.ApiModel
	err = global.DB.Find(&apiModels).Error
	if err != nil {
		return nil, nil, fmt.Errorf("GetElTreeApis: find -> %v", err)
	}

	var apiGroup []string
	err = global.DB.Model(&systemModel.ApiModel{}).Distinct().Pluck("api_group", &apiGroup).Error
	if err != nil {
		return nil, nil, fmt.Errorf("GetElTreeApis: apiGroup -> %v", err)
	}

	// 前端 el-tree data
	treeData := make(map[string][]systemModel.Children, len(apiModels))
	for _, model := range apiModels {
		var children systemModel.Children
		sPath := strings.Split(model.Path, fmt.Sprintf("%s/", model.ApiGroup))
		var tPath string
		if len(sPath) == 2 {
			tPath = sPath[1]
		}
		children.Key = fmt.Sprintf("%s,%s", model.Path, model.Method)
		children.ApiGroup = fmt.Sprintf("%s -> %s", tPath, model.Description)
		children.Path = model.Path
		children.Method = model.Method
		children.Description = model.Description
		treeData[model.ApiGroup] = append(treeData[model.ApiGroup], children)
	}

	for _, value := range apiGroup {
		var apiTree systemModel.ApiTree
		apiTree.ApiGroup = value
		apiTree.Children = treeData[value]
		list = append(list, apiTree)
	}

	// 前端 el-tree default-checked-keys
	e := CasbinServiceApp.Casbin()
	authorityId := strconv.Itoa(int(roleId))
	cData := e.GetFilteredPolicy(0, authorityId)
	for _, v := range cData {
		checkedKey = append(checkedKey, fmt.Sprintf("%s,%s", v[1], v[2]))
	}

	return
}

// DeleteApi 删除指定api
func (a *ApiService) DeleteApi(id uint) (err error) {
	var apiModel systemModel.ApiModel
	err = global.DB.Where("id = ?", id).First(&apiModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOG.Error("deleteApi -> 查找id", zap.Error(err))
		return err
	}

	err = global.DB.Unscoped().Delete(&apiModel).Error
	if err != nil {
		global.LOG.Error("deleteApi -> 删除id", zap.Error(err))
		return err
	}

	ok := CasbinServiceApp.ClearCasbin(1, apiModel.Path, apiModel.Method)
	if !ok {
		global.LOG.Warn("ApiPath: " + apiModel.Path + ",Method: " + apiModel.Method + " casbin同步清理失败")
	}
	e := CasbinServiceApp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

// EditApi 编辑api
func (a *ApiService) EditApi(eApi systemReq.EditApi) (err error) {
	var oldApiModel systemModel.ApiModel
	err = global.DB.Where("id = ?", eApi.Id).First(&oldApiModel).Error
	if err != nil {
		return errors.New("editApi: id不存在")
	}

	if oldApiModel.Path != eApi.Path || oldApiModel.Method != eApi.Method {
		if !errors.Is(global.DB.Where("path = ? AND method = ?", eApi.Path, eApi.Method).First(&systemModel.ApiModel{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("editApi: 存在相同接口")
		}
	}

	err = CasbinServiceApp.UpdateCasbinApi(oldApiModel.Path, eApi.Path, oldApiModel.Method, eApi.Method)
	if err != nil {
		return fmt.Errorf("editApi: 更新casbin rule -> %v", err)
	}

	return global.DB.Debug().Model(&oldApiModel).Updates(map[string]interface{}{"path": eApi.Path, "method": eApi.Method, "api_group": eApi.ApiGroup, "description": eApi.Description}).Error
}

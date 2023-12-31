package system

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
)

type UserService struct{}

// Login 登陆校验
func (us *UserService) Login(u *systemModel.UserModel) (userInter *systemModel.UserModel, err error) {
	var userModel systemModel.UserModel
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	if err != nil {
		return nil, errors.New("用户不存在或密码不正确")
	}
	if userModel.Active == false {
		return nil, errors.New("用户为禁用状态")
	}
	return &userModel, err
}

func (us *UserService) GetUserInfo(userId uint) (userResults systemRes.UserResult, err error) {
	err = global.DB.Table("sys_user").Select("sys_user.created_at,sys_user.id,sys_user.username,sys_user.phone,sys_user.email,sys_user.active,sys_user.role_model_id,sys_role.role_name").Joins("inner join sys_role on sys_user.role_model_id = sys_role.id").Where("sys_user.id = ?", userId).Scan(&userResults).Error
	return
}

// GetUsers 获取所有用户
func (us *UserService) GetUsers(userSp systemReq.UserSearchParams) ([]systemRes.UserResult, int64, error) {
	var userResults []systemRes.UserResult
	var total int64

	db := global.DB.Model(&systemModel.UserModel{})

	if userSp.Name != "" {
		db = db.Where("username LIKE ?", "%"+userSp.Name+"%")
	}

	// 分页
	err := db.Count(&total).Error
	if err != nil {
		return userResults, total, fmt.Errorf("分页count -> %v", err)
	} else {
		limit := userSp.PageSize
		offset := userSp.PageSize * (userSp.Page - 1)
		db = db.Limit(limit).Offset(offset)
		//err = db.Find(&list).Error
		// 左连接 查询出role_name
		db.Select("sys_user.id,sys_user.username,sys_user.phone,sys_user.email,sys_user.active,sys_user.role_model_id,sys_role.role_name").Joins("left join sys_role on sys_user.role_model_id = sys_role.id").Scan(&userResults)
	}

	return userResults, total, err
}

// DeleteUser 删除用户
func (us *UserService) DeleteUser(id uint) (err error) {
	return global.DB.Where("id = ?", id).Unscoped().Delete(&systemModel.UserModel{}).Error
}

// AddUser 添加用户
func (us *UserService) AddUser(user systemReq.AddUser) (err error) {
	err = global.DB.Where("id = ?", user.RoleModelID).First(&systemModel.RoleModel{}).Error
	if err != nil {
		global.LOG.Error("添加用户 -> 查询role", zap.Error(err))
		return err
	}

	var userModel systemModel.UserModel
	userModel.Username = user.Username
	userModel.Password = utils.MD5V([]byte(user.Password))
	userModel.Phone = user.Phone
	userModel.Email = user.Email
	userModel.Active = user.Active
	userModel.RoleModelID = user.RoleModelID

	return global.DB.Create(&userModel).Error
}

// EditUser 编辑用户
func (us *UserService) EditUser(user systemReq.EditUser) (*systemRes.UserResult, error) {
	var userModel systemModel.UserModel
	var userResult systemRes.UserResult
	// 用户是否存在
	err := global.DB.Where("id = ?", user.Id).First(&userModel).Error
	if err != nil {
		global.LOG.Error("编辑用户 -> 查询Id", zap.Error(err))
		return nil, err
	}

	// 角色是否存在
	var roleModel systemModel.RoleModel
	err = global.DB.Where("id = ?", user.RoleModelID).First(&roleModel).Error
	if err != nil {
		global.LOG.Error("编辑用户 -> 查询role", zap.Error(err))
		return nil, err
	}

	updateV := make(map[string]interface{}, 5)
	updateV["username"] = user.Username
	updateV["active"] = user.Active
	updateV["role_model_id"] = user.RoleModelID
	updateV["phone"] = user.Phone
	updateV["email"] = user.Email

	err = global.DB.Model(&userModel).Updates(updateV).Error
	if err != nil {
		global.LOG.Error("编辑用户 -> update", zap.Error(err))
		return nil, err
	}

	userResult.ID = userModel.ID
	userResult.Username = userModel.Username
	userResult.Phone = userModel.Phone
	userResult.Email = userModel.Email
	userResult.Active = userModel.Active
	userResult.RoleName = roleModel.RoleName
	userResult.RoleModelID = userModel.RoleModelID

	return &userResult, nil
}

// ModifyPass 修改用户密码
func (us *UserService) ModifyPass(mp systemReq.ModifyPass) (err error) {
	var userModel systemModel.UserModel
	err = global.DB.Where("id = ? and password = ?", mp.Id, utils.MD5V([]byte(mp.OldPassword))).First(&userModel).Error
	if err != nil {
		global.LOG.Error("修改用户密码 -> 查询用户", zap.Error(err))
		return err
	}
	return global.DB.Model(&userModel).Update("password", utils.MD5V([]byte(mp.NewPassword))).Error
}

// SwitchActive 切换启用状态
func (us *UserService) SwitchActive(sa systemReq.SwitchActive) (err error) {
	var userModel systemModel.UserModel
	err = global.DB.Where("id = ?", sa.Id).First(&userModel).Error
	if err != nil {
		global.LOG.Error("切换启用状态 -> 查询用户", zap.Error(err))
		return err
	}
	if sa.Active {
		return global.DB.Model(&userModel).Update("active", true).Error
	} else {
		return global.DB.Model(&userModel).Update("active", false).Error
	}
}

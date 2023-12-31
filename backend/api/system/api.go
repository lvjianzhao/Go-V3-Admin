package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	systemRep "server/model/system/response"
	"server/utils"
	"strconv"
)

type ApiApi struct{}

// table  测试数据
func (a *ApiApi) GetTable(c *gin.Context) {
	list := []map[string]interface{}{
		{
			"id":         "630000201510116854",
			"username":   "Shirley Anderson",
			"phone":      "18575455423",
			"email":      "s.kdqjzhvbt@sjopf.sl",
			"roles":      "admin",
			"status":     true,
			"createTime": "1983-03-12 15:49:10",
		},
		{
			"id":         "440000200009077556",
			"username":   "Helen Martin",
			"phone":      "13805108397",
			"email":      "h.pofemkqqyt@wotbxcx.tt",
			"roles":      "editor",
			"status":     false,
			"createTime": "1995-08-28 04:13:18",
		},
		{
			"id":         "630000201510116854",
			"username":   "Shirley Anderson",
			"phone":      "18575455423",
			"email":      "s.kdqjzhvbt@sjopf.sl",
			"roles":      "admin",
			"status":     true,
			"createTime": "1983-03-12 15:49:10",
		},
		{
			"id":         "440000200009077556",
			"username":   "Helen Martin",
			"phone":      "13805108397",
			"email":      "h.pofemkqqyt@wotbxcx.tt",
			"roles":      "editor",
			"status":     false,
			"createTime": "1995-08-28 04:13:18",
		},
		{
			"id":         "630000201510116854",
			"username":   "Shirley Anderson",
			"phone":      "18575455423",
			"email":      "s.kdqjzhvbt@sjopf.sl",
			"roles":      "admin",
			"status":     true,
			"createTime": "1983-03-12 15:49:10",
		},
		{
			"id":         "440000200009077556",
			"username":   "Helen Martin",
			"phone":      "13805108397",
			"email":      "h.pofemkqqyt@wotbxcx.tt",
			"roles":      "editor",
			"status":     false,
			"createTime": "1995-08-28 04:13:18",
		},
		{
			"id":         "630000201510116854",
			"username":   "Shirley Anderson",
			"phone":      "18575455423",
			"email":      "s.kdqjzhvbt@sjopf.sl",
			"roles":      "admin",
			"status":     true,
			"createTime": "1983-03-12 15:49:10",
		},
		{
			"id":         "440000200009077556",
			"username":   "Helen Martin",
			"phone":      "13805108397",
			"email":      "h.pofemkqqyt@wotbxcx.tt",
			"roles":      "editor",
			"status":     false,
			"createTime": "1995-08-28 04:13:18",
		},
		{
			"id":         "630000201510116854",
			"username":   "Shirley Anderson",
			"phone":      "18575455423",
			"email":      "s.kdqjzhvbt@sjopf.sl",
			"roles":      "admin",
			"status":     true,
			"createTime": "1983-03-12 15:49:10",
		},
		{
			"id":         "440000200009077556",
			"username":   "Helen Martin",
			"phone":      "13805108397",
			"email":      "h.pofemkqqyt@wotbxcx.tt",
			"roles":      "editor",
			"status":     false,
			"createTime": "1995-08-28 04:13:18",
		},
	}

	apiSp := struct {
		Page     int
		PageSize int
	}{Page: 1, PageSize: 10}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    10,
		Page:     apiSp.Page,
		PageSize: apiSp.PageSize,
	}, "获取成功", c)

}

// AddApi 添加api
func (a *ApiApi) AddApi(c *gin.Context) {
	var apiReq systemModel.ApiModel
	_ = c.ShouldBindJSON(&apiReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if instance, err := apiService.AddApi(apiReq); err != nil {
		response.FailWithMessage("添加失败", c)
		global.LOG.Error("添加失败", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "添加成功", c)
	}
}

// GetApis 列出所有/根据条件搜索api
func (a *ApiApi) GetApis(c *gin.Context) {
	var apiSp systemReq.ApiSearchParams
	apiSp.ApiGroup = c.Query("api_group")
	apiSp.Path = c.Query("path")
	apiSp.Method = c.Query("method")
	apiSp.Description = c.Query("description")
	apiSp.ApiGroup = c.Query("api_group")
	apiSp.Page, _ = utils.StringToInt(c.Query("page"))
	apiSp.PageSize, _ = utils.StringToInt(c.Query("pageSize"))
	apiSp.OrderKey = c.Query("orderKey")
	apiSp.Desc, _ = strconv.ParseBool(c.Query("desc"))

	if list, total, err := apiService.GetApis(apiSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     apiSp.Page,
			PageSize: apiSp.PageSize,
		}, "获取成功", c)
	}
}

// GetApiGroups 获取所有API分组
func (a *ApiApi) GetApiGroups(c *gin.Context) {
	if list, total, err := apiService.GetApiGroups(); err != nil {
		response.FailWithMessage("获取失败", c)
		global.LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// DeleteApi 删除指定api
func (a *ApiApi) DeleteApi(c *gin.Context) {
	var cId request.CId
	cId.ID, _ = utils.StringToUint(c.Param("apiID"))

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.DeleteApi(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// EditApi 编辑api
func (a *ApiApi) EditApi(c *gin.Context) {
	var eApi systemReq.EditApi
	_ = c.ShouldBindJSON(&eApi)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&eApi); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.EditApi(eApi); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// GetElTreeApis 格式化列出所有api
func (a *ApiApi) GetElTreeApis(c *gin.Context) {
	var cId request.CId
	cId.ID, _ = utils.StringToUint(c.Query("id"))

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	list, checkedKey, err := apiService.GetElTreeApis(cId.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRep.ApiTree{
			List:       list,
			CheckedKey: checkedKey,
		}, "获取成功", c)
	}
}

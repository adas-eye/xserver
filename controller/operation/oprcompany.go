package operation

import (
	"xserver/middleware"
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Company
type Company struct {
}

// ListHandler 列表
func (o *Company) ListHandler(c *gin.Context) {
	var p orm.DbPage
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	where := p.DbWhere()
	where.Append("parent_id = ?", 0) // 上级节点为0，表示公司
	where.Append("guid != ?", "")
	var data []model.OprOrganization
	toatl, _ := orm.DbByWhere(&model.OprOrganization{}, where).Find(&data)
	ctx.JSONWrite(gin.H{"total": toatl, "data": data}, c)
}

// GetHandler 详细
func (o *Company) GetHandler(c *gin.Context) {
	service.QueryById(&model.OprOrganization{}, c)
}

// AddHandler 新增
func (o *Company) AddHandler(c *gin.Context) {
	var p model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&p.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	p.GUID = util.NUID()
	tok := middleware.GetUserToken(c)
	u := &model.SysUser{}
	u.UserName = p.UserName
	u.CreatedBy = tok.UserName
	u.OrganizeName = p.Name
	u.Type = model.SysUserTypeAdmin
	u.OrganizeGUID = p.GUID
	u.DeviceIds = "*"
	if err := service.SysUserCreate(u); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *Company) UpdateHandler(c *gin.Context) {
	var p model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&p.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *Company) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.OprOrganization{}, c)
}

func CompanyRouters(r *gin.RouterGroup) {
	o := Company{}
	r.GET("/list", o.ListHandler)
	r.GET("/:id", o.GetHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}

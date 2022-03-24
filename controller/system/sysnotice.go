package system

import (
	"xserver/model"
	"xserver/service"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Notice
type Notice struct {
}

// ListHandler 列表
func (o *Notice) ListHandler(c *gin.Context) {
	var param service.BasePage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysNotice
	total, _ := orm.DbPage(&model.SysNotice{}, param.Where()).Find(param.PageNum, param.PageSize, &data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 详细
func (o *Notice) GetHandler(c *gin.Context) {
	var data model.SysNotice
	service.QueryById(&data, c)
}

// AddHandler 新增
func (o *Notice) AddHandler(c *gin.Context) {
	var data model.SysNotice
	//获取参数
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Notice) UpdateHandler(c *gin.Context) {
	var data model.SysNotice
	//获取参数
	if err := c.ShouldBind(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Notice) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysNotice{}, c)
}

func NoticeRouters(r *gin.RouterGroup) {
	sysNotice := Notice{}
	r.GET("/notice/list", sysNotice.ListHandler)
	r.GET("/notice/:id", sysNotice.GetHandler)
	r.POST("/notice", sysNotice.AddHandler)
	r.PUT("/notice", sysNotice.UpdateHandler)
	r.DELETE("/notice/:id", sysNotice.DeleteHandler)
}
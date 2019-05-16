package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	CommonController
}

func NewUserController() *UsersController {
	o := &UsersController{}
	return o
}
func (this UsersController) Copy() *UsersController { //
	cp := this
	this.ctx = nil
	return &cp
}
func (this UsersController) Add(ctx *gin.Context) { //
	userinfo := new(models.Users)
	if err := ctx.ShouldBind(&userinfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := userinfo.Add()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(id)
}

// func (this UsersController) Inquire(ctx *gin.Context) { //
// 	userdata := new(models.Users)

// 	_userinfo, err := userdata.Inquire()
// 	if err != nil {
// 		this.RenderError(err)
// 		return
// 	}
// 	this.Render(_userinfo)
// }

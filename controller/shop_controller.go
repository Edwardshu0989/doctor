package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type ShopsController struct {
	CommonController
}

func NewShopController() *ShopsController {
	o := &ShopsController{}
	return o
}
func (this ShopsController) Copy() *ShopsController { //
	cp := this
	this.ctx = nil
	return &cp
}
func (this ShopsController) Add(ctx *gin.Context) { //
	shopinfo := new(models.Shop)
	if err := ctx.ShouldBind(&shopinfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := shopinfo.Add()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(id)
}
func (this ShopsController) Inquire(ctx *gin.Context) { //
	shopinfo := new(models.Shop)

	_shopdata, err := shopinfo.Inquire()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(_shopdata)
}

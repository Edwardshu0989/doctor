package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	CommonController
}

func NewOrderController() *OrderController {
	o := &OrderController{}
	return o
}

func (this OrderController) Copy() *OrderController { //
	cp := this
	this.ctx = nil
	return &cp
}

func (this OrderController) Add(ctx *gin.Context) { //
	medicalinfo := new(models.Medical)
	if err := ctx.ShouldBind(&medicalinfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := medicalinfo.Add()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(id)

}

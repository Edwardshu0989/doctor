package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	CommonController
}

func NewProductController() *ProductsController {
	o := &ProductsController{}
	return o
}

func (this ProductsController) Copy() *ProductsController { //
	cp := this
	this.ctx = nil
	return &cp
}

func (this ProductsController) Add(ctx *gin.Context) { //
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
func (this ProductsController) Inquire(ctx *gin.Context) { //
	medicalinfo := new(models.Medical)

	_medicaldata, err := medicalinfo.Inquire()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(_medicaldata)
}

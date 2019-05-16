package controller

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type TechniciansController struct {
	CommonController
}

func NewTechnicianController() *TechniciansController {
	o := &TechniciansController{}
	return o
}
func (this TechniciansController) Copy() *TechniciansController { //
	cp := this
	this.ctx = nil
	return &cp
}
func (this TechniciansController) Add(ctx *gin.Context) { //
	technicianinfo := new(models.Technician)
	if err := ctx.ShouldBind(&technicianinfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := technicianinfo.Add()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(id)

}
func (this TechniciansController) Inquire(ctx *gin.Context) { //
	technicianinfo := new(models.Technician)

	_technicianinfo, err := technicianinfo.Inquire()
	if err != nil {
		this.RenderError(err)
		return
	}
	this.Render(_technicianinfo)
}

package controller

import (
	"github.com/gin-gonic/gin"
)

var (
	Ginc *gin.Context
)

type HttpStatus int

const (
	HttpStatusOk             HttpStatus = 200
	HttpStatusAuthErr        HttpStatus = 401
	HttpStatusLoginErr       HttpStatus = 402
	HttpStatusAlreadyExist   HttpStatus = 403
	HttpStatusNotBelong      HttpStatus = 404
	HttpStatusParamsErr      HttpStatus = 500
	HttpStatusServerErr      HttpStatus = 501
	HttpStatusEntityTooLarge HttpStatus = 502
)

type CommonController struct {
	ctx *gin.Context
}

func (this *CommonController) RenderError(data interface{}) {
	h := gin.H{
		"code": 501,
		"data": data}
	Ginc.JSON(501, h)
}
func (this *CommonController) Render(data interface{}) {
	h := gin.H{
		"code": 200,
		"data": data}
	Ginc.JSON(200, h)

}

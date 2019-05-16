package main

import (
	"./controller"
	"./libs/middleware"
	"github.com/gin-gonic/gin"
)

func Newserver() (r *gin.Engine, err error) {
	r = gin.Default()
	r.Use(middleware.Before)
	v1 := r.Group("/shop")
	{
		controller := controller.NewShopController() //同包名下调用方法（方法的首字母大写）
		v1.POST("/add", controller.Copy().Add)
		v1.GET("/inquire", controller.Copy().Inquire)

	}
	v2 := r.Group("/product")
	{
		controller := controller.NewProductController() //同包名下调用方法（方法的首字母大写）
		v2.POST("/add", controller.Copy().Add)
		v2.GET("/inquire", controller.Copy().Inquire)

	}
	v3 := r.Group("/technician")
	{
		controller := controller.NewTechnicianController() //同包名下调用方法（方法的首字母大写）
		v3.POST("/add", controller.Copy().Add)
		v3.GET("/inquire", controller.Copy().Inquire)

	}
	v4 := r.Group("/user")
	{
		controller := controller.NewUserController() //同包名下调用方法（方法的首字母大写）
		v4.POST("/add", controller.Copy().Add)
		// v3.GET("/inquire", controller.Copy().Inquire)

	}
	v5 := r.Group("/order")
	{
		controller := controller.NewOrderController() //同包名下调用方法（方法的首字母大写）
		v5.POST("/add", controller.Copy().Add)
		// v3.GET("/inquire", controller.Copy().Inquire)

	}
	return
}

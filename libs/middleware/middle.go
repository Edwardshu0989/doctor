package middleware

import (
	"../../controller"
	"github.com/gin-gonic/gin"
)

func Before(g *gin.Context) {
	controller.Ginc = g

	g.Next()
}

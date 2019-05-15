package api1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/api1")
	route.GET("", GET)
	route.POST("", POST)

}

func GET(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"route": "/apiGroup1/api1"})
}

func POST(ctx *gin.Context) {

}

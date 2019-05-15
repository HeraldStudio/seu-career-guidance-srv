package apiGroup1

import (
	"github.com/gin-gonic/gin"
	"go-common/handlers/apiGroup1/api1"
	"net/http"
)

func Handler(router *gin.RouterGroup) {

	route := router.Group("/apiGroup1")
	route.GET("", GET)
	route.POST("", POST)
	// 注册下一层 handler
	api1.Handler(router)

}

func GET(ctx *gin.Context) {
	// 处理 GET - /apiGroup1 请求
	ctx.JSON(http.StatusOK, gin.H{"route": "/apiGroup1"})
}

func POST(ctx *gin.Context) {
	// 处理 POST - /apiGroup1 请求
}

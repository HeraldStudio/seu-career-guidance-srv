package handlers

import (
	"github.com/gin-gonic/gin"
	"go-common/handlers/apiGroup1"
)

func Handler(engine *gin.Engine) {

	router := engine.Group("/")

	// 注册所有顶层handler
	apiGroup1.Handler(router)

}

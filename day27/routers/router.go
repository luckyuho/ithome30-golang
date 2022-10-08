package routers

import (
	v1 "it/day27/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 定義前端打的 api 路徑
	r.GET("/v1/hello_world", v1.HelloWorld)
	r.POST("/v1/register", v1.ApiRegister)
	r.POST("/v1/login", v1.ApiLogin)
	return r
}

package routers

import (
	v1 "it/day24/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// 定義前端打的 api 路徑
	r.GET("/v1/hello_world", v1.HelloWorld)
	return r
}

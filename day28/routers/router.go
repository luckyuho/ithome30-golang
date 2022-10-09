package routers

import (
	mJwt "it/day28/app/middleware/jwt"
	v1 "it/day28/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 定義前端打的 api 路徑
	r.POST("/v1/register", v1.ApiRegister)
	r.POST("/v1/login", v1.ApiLogin)

	// 需要用到 jwt 的打包在一起
	user := r.Group("/v1/user")
	user.Use(mJwt.JWTAuth())
	user.GET("/info", v1.ApiGetUserInfo)

	return r
}

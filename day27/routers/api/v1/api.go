package v1

import (
	"net/http"

	ctr "it/day27/app/controllers/user"
	api "it/day27/routers/api"

	"github.com/gin-gonic/gin"
)

// 如果是要給前端的 api，函式只能有一個輸入參數 *gin.Context 且不能有輸出
func HelloWorld(c *gin.Context) {
	// 包成 json 的格式丟給前端
	c.JSON(http.StatusOK, gin.H{
		"data": "Hello world!",
	})
}

// 註冊使用者
func ApiRegister(c *gin.Context) {
	input := api.User{}
	c.Bind(&input)
	result := ctr.RegisterUser(input.Email, input.Password)
	template(c, http.StatusOK, result)
}

// 使用者登入
func ApiLogin(c *gin.Context) {
	input := api.User{}
	c.Bind(&input)
	result := ctr.LoginUser(input.Email, input.Password)
	template(c, http.StatusOK, result)
}

// 固定形式的輸出，有利於前端使用
func template(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"data": data,
	})
}

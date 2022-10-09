package v1

import (
	"net/http"

	ctr "it/day28/app/controllers/user"
	mJwt "it/day28/app/middleware/jwt"
	t "it/day28/app/utils/template"
	api "it/day28/routers/api"

	"github.com/gin-gonic/gin"
)

// 註冊使用者
func ApiRegister(c *gin.Context) {
	input := api.User{}
	c.Bind(&input)
	result := ctr.RegisterUser(input.Email, input.Password)
	t.Template(c, http.StatusOK, result)
}

// 使用者登入
func ApiLogin(c *gin.Context) {
	input := api.User{}
	c.Bind(&input)
	result := ctr.LoginUser(input.Email, input.Password)
	t.Template(c, http.StatusOK, result)
}

// 獲取 jwt token 中的資訊
func ApiGetUserInfo(c *gin.Context) {
	claim, err := mJwt.GetUserInfo(c)
	if err != nil {
		t.Template(c, http.StatusBadRequest, err)
	}
	t.Template(c, http.StatusOK, claim)
}

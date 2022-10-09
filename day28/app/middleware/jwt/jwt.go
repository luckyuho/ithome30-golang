package jwt

import (
	"fmt"
	"net/http"
	"os"
	"time"

	t "it/day28/app/utils/template"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv/autoload"
)

var Secret = os.Getenv("SECRET")
var Issuer = os.Getenv("Issuer")

type AuthClaims struct {
	jwt.StandardClaims
	Email string
	Id    int
}

// 解析 jwt
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通過 http header 中的 token 解析認證
		token := c.Request.Header.Get("token")
		if token == "" {
			t.Template(c, http.StatusBadRequest, "not jwt token")
			c.Abort()
			return
		}

		// 解析 jwt 是否正確，如果不正確則提前結束，正確就繼續
		_, err := parseToken(token)
		if err != nil {
			var errMsg string
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errMsg = "token無效"
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					errMsg = "token過期"
				}
			}
			t.Template(c, http.StatusBadRequest, errMsg)
			c.Abort()
			return
		}
	}
}

// 獲取使用者資訊
func GetUserInfo(c *gin.Context) (*AuthClaims, error) {
	// 通過http header中的token解析來認證
	token := c.Request.Header.Get("token")
	if token == "" {
		return nil, fmt.Errorf("no jwt token")
	}

	claim, err := parseToken(token)
	if err != nil {
		return nil, fmt.Errorf("bad jwt: %s", err)
	}

	return claim, nil
}

// 解析 jwt token
func parseToken(token string) (*AuthClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*AuthClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

// 產生 jwt
func GenerateJWT(id int, email string) (string, error) {
	expiresAt := time.Now().Add(10 * time.Second).Unix()
	claims := AuthClaims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer:    Issuer,
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

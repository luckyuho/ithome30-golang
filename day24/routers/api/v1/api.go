package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 如果是要給前端的 api，函式只能有一個輸入參數 *gin.Context 且不能有輸出
func HelloWorld(c *gin.Context) {
	// 包成 json 的格式丟給前端
	c.JSON(http.StatusOK, gin.H{
		"data": "Hello world!",
	})
}

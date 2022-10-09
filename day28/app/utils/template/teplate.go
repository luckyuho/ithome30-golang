package template

import "github.com/gin-gonic/gin"

// 固定形式的輸出，有利於前端使用
func Template(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"data": data,
	})
}

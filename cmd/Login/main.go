package main

import (
	"LaunchService/LaunchUser/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	var form Login
	r.POST("/loginForm", func(c *gin.Context) {
		// 声明接收的变量
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		correct, answer := client.LaunchLoginClient(form.User, form.Pssword)
		if correct == true {
			c.JSON(http.StatusOK, gin.H{"status": "200"})
			c.String(200, answer)
		}
		if correct == false {
			c.String(404, answer)
		}

	})
	r.Run(":8000")
}

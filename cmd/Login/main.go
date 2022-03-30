package main

import (
	"launchservice/launchuser/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	var form login
	r.POST("/loginForm", func(c *gin.Context) {
		// 声明接收的变量
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		nameCorrect, PasswordCorrect := client.LaunchLoginClient(form.User, form.Pssword)
		if nameCorrect == true && PasswordCorrect == true {
			//c.JSON(http.StatusOK, gin.H{"status": "200"})
			c.String(200, "登陆成功")
		}
		if nameCorrect == false {
			c.String(201, "账号不存在")
		}
		if PasswordCorrect == false {
			c.String(202, "密码错误")
		}
	})
	r.Run(":8000")
}

package loginservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func GinLogin() (username string, password string) {
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
		// // 判断用户名密码是否正确
		// if form.User != "root" || form.Pssword != "admin" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
	return form.User, form.Pssword
}

func GinHello(username string, greeting string) {
	route := gin.Default()
	route.POST("/loginForm", func(context *gin.Context) {
		context.String(http.StatusOK, "welcome "+username+"\n"+greeting)
	})
	route.Run(":8001")
}

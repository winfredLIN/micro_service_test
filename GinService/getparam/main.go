package getparam

import (
	"net/http"
	//"strings"

	"github.com/gin-gonic/gin"
)

func Greetinghandler(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "hello "+name)
}

func Getparam() (name string) {
	route := gin.Default()
	route.GET("/user/:name", Greetinghandler)
	route.Run(":8080")
	return name
}

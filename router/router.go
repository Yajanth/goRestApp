package router

import (
	"github.com/Yajanth/goRestApp/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", handler.GetUsers)

	return r
}

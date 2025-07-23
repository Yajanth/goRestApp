package handler

import (
	"net/http"

	"github.com/Yajanth/goRestApp/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

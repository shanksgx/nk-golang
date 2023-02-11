package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello golang",
	})
}

func CreateHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello golang post",
	})
}

package handlers

import "github.com/gin-gonic/gin"

type UserHandlerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

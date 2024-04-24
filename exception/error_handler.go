package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Error(err.(error))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}

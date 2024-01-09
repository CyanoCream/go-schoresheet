package middleware

import (
	"github.com/gin-gonic/gin"
	"go-scoresheet/master/helpers"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

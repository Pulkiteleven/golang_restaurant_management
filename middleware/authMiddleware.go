package middleware

import (
	"fmt"
	"net/http"
	"restaurant_management/helpers"

	"github.com/gin-gonic/gin"
)

func Authetication()  gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":fmt.Sprintf("No Authorization Provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			c.Abort()
			return
		}
		c.Set("email",claims.Email)
		c.Set("First_name",claims.First_name)
		c.Set("Last_name",claims.Last_name)
		c.Set("uid",claims.Uid)

		c.Next()

	}
}
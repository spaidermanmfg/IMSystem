package middlewares

import (
	"IMSystem/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Check the user login status.
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := tools.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "用户认证失败",
			})
			return
		}
		c.Set("userClaims", userClaims)
		c.Next()
	}

}

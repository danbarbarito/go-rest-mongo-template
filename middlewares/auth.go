package middlewares

import (
	"github.com/gin-gonic/gin"
	"myapi/models"
	db "myapi/models/db"
	"myapi/services"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Bearer-Token")
		tokenModel, err := services.VerifyToken(token, db.TokenTypeAccess)
		if err != nil {
			models.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("userIdHex", tokenModel.User.Hex())
		c.Set("userId", tokenModel.User)

		c.Next()
	}
}

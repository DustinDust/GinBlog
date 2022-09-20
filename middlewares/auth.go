package middlewares

import (
	"net/http"
	"strings"

	"github.com/DustinDust/gin-blog-post/controllers"
	"github.com/DustinDust/gin-blog-post/models"
	"github.com/DustinDust/gin-blog-post/services"
	"github.com/gin-gonic/gin"
)

func JwtRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerTokenString := ctx.GetHeader("Authorization")
		claims, err := services.JwtService.ParseJwtFromBearerToken(bearerTokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		ctx.Set("userId", claims.UserId)
		ctx.Next()
	}
}

func RefreshTokenRequiredMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerTokenString := c.GetHeader("Authorization")
		tokenString := strings.Fields(bearerTokenString)[1]
		claims, err := services.JwtService.ParseJwtFromBearerToken(bearerTokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		userId := claims.UserId
		res, user := models.UserRepository.FindById(userId)
		if res.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		if user.RefreshToken != tokenString {
			c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Refresh token doens't match",
			})
			return
		}
		c.Set("userId", userId)
		c.Set("refreshToken", tokenString)
		c.Next()
	}
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/DustinDust/gin-blog-post/models"
	"github.com/DustinDust/gin-blog-post/services"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type authController struct{}

func (a *authController) JwtTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"UserId": c.MustGet("userId"),
	})
}

func (a *authController) Login(c *gin.Context) {
	loginDto := &LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res, user := models.UserRepository.FindByUsername(loginDto.Username)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	if user.Password == loginDto.Password {
		accessToken, err1 := services.JwtService.GenerateJwt(services.JwtClaims{
			UserId: int(user.ID),
		})
		refreshToken, err2 := services.JwtService.GenerateRefreshJwt(services.JwtClaims{
			UserId: int(user.ID),
		})
		models.UserRepository.Update(int(user.ID), models.User{RefreshToken: refreshToken})
		if err1 != nil || err2 != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: fmt.Sprintf("[%v, %v]", err1, err2),
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "OK",
			Data: gin.H{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "Wrong password",
		})
		return
	}
}

func (a *authController) Register(c *gin.Context) {
	loginDto := LoginDto{}
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res, user := models.UserRepository.Create(models.User{
		Username: loginDto.Username,
		Password: loginDto.Password,
	})
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data:    user,
	})
}

func (a *authController) Refresh(c *gin.Context) {
	userId, ok := c.MustGet("userId").(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "Invalid token",
		})
		return
	}
	refreshtokenString := c.MustGet("refreshToken").(string)
	token, err := services.JwtService.GenerateJwt(services.JwtClaims{
		UserId: userId,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data: gin.H{
			"access_token":  token,
			"refresh_token": refreshtokenString,
		},
	})
}

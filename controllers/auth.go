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

// Login godoc
// @Summary Login
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param        login body      controllers.LoginDto true  "Login"
// @Success 200 {object} controllers.Response{data=controllers.Abr{accessToken=string,refreshToken=string}}
// @Failure 401 {object} controllers.Response
// @Failure 500 {object} controllers.Response
// @Router /v1/auth/login [post]
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
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
		if err1 != nil || err2 != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: fmt.Sprintf("[%v, %v]", err1, err2),
			})
			return
		}
		models.UserRepository.Update(int(user.ID), models.User{RefreshToken: refreshToken})
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

// Register godoc
// @Summary Register
// @Description Register with username and password
// @Tags auth
// @Param Register body controllers.LoginDto true "Register"
// @Accept json
// @Produce json
// @Success 200 {object} controllers.Response{data=controllers.Abr{accessToken=string,refreshToken=string}}
// @Router /v1/auth/resgiter [post]
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

// Refresh godoc
// @Summary Refresh access token
// @Description Refresh access token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} controllers.Response{data=controllers.Abr{accessToken=string,refreshToken=string}}
// @Security ApiKeyuAuth
// @Router /v1/auth/refresh [post]
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

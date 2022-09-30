package controllers

import (
	"net/http"

	"github.com/DustinDust/gin-blog-post/models"
	"github.com/gin-gonic/gin"
)

type userController struct{}
type CreateUserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserDto struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Profile godoc
// @Summary Get current user profile
// @Description Base on userId extracted from jwt token, get user profile
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} controllers.Response{data=models.User}
// @Security ApiKeyuAuth
// @Router /v1/user/me [get]
func (u *userController) FindMe(c *gin.Context) {
	id, ok := c.MustGet("userId").(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Authorization failed",
		})
		return
	}
	res, user := models.UserRepository.FindById(id)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Success",
			Data:    user,
		})
	}
}

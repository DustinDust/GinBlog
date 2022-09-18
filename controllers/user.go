package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/DustinDust/gin-blog-post/models"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface struct{}
type CreateUserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserDto struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *UserControllerInterface) FindById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
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

func (u *UserControllerInterface) Create(c *gin.Context) {
	createDto := CreateUserDto{}
	err := c.ShouldBindJSON(&createDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	user := models.User{
		Username: createDto.Username,
		Password: createDto.Password,
	}
	res, user := models.UserRepository.Create(user)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Success",
		Data:    user,
	})
}

func (u *UserControllerInterface) Update(c *gin.Context) {
	updateDto := UpdateUserDto{}
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("Error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	user := models.User{
		Username: updateDto.Username,
		Password: updateDto.Password,
	}
	res := models.UserRepository.Update(id, user)
	if res.Error != nil {
		log.Printf("Error: %v", res.Error.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data: gin.H{
			"rowAffected": res.RowsAffected,
		},
	})
}

func (u *UserControllerInterface) Delete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res := models.UserRepository.Delete(id)
	if res.Error != nil {
		if res.Error != nil {
			log.Printf("Error: %v", res.Error.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: res.Error.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Success",
		Data: gin.H{
			"rowAffected": res.RowsAffected,
		},
	})
}

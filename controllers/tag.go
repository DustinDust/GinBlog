package controllers

import (
	"net/http"
	"strconv"

	"github.com/DustinDust/gin-blog-post/models"
	"github.com/gin-gonic/gin"
)

type TagControllerInterface struct{}

type CreateTagDto struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type UpdateTagDto struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

func (t *TagControllerInterface) FindAll(c *gin.Context) {
	query := make(map[string]interface{})
	urlQueries := c.Request.URL.Query()

	for key, elem := range urlQueries {
		query[key] = elem
	}

	res, tags := models.TagRepository.FindAll(query)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data:    tags,
	})
}

func (t *TagControllerInterface) FindById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res, tag := models.TagRepository.FindById(id)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data:    tag,
	})
}

func (t *TagControllerInterface) Create(c *gin.Context) {
	createDto := CreateTagDto{}
	if err := c.ShouldBindJSON(&createDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	tag := models.Tag{
		Name: createDto.Name,
		Code: createDto.Code,
	}
	res, tag := models.TagRepository.Create(tag)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data:    tag,
	})
}

func (t *TagControllerInterface) Update(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	updateDto := UpdateTagDto{}
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	tag := models.Tag{
		Name: updateDto.Name,
		Code: updateDto.Code,
	}
	res := models.TagRepository.Update(id, tag)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data: gin.H{
			"rowsAffected": res.RowsAffected,
		},
	})
}

func (t *TagControllerInterface) Delete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res := models.TagRepository.Delete(id)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "OK",
		Data: gin.H{
			"rowsAffected": res.RowsAffected,
		},
	})
}

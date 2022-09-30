package controllers

import (
	"net/http"
	"strconv"

	"github.com/DustinDust/gin-blog-post/models"
	"github.com/gin-gonic/gin"
)

type tagController struct{}

type CreateTagDto struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type UpdateTagDto struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

// "Get All" godoc
// @Summary Get All Tags
// @Description Get Tags with query and paging
// @Tags tag
// @Produce json
// @Param page query integer true "page index"
// @Success 200 {object} controllers.Response{data=models.Pagination[models.Tag]}
// @Router /v1/tag [get]
func (t *tagController) FindAll(c *gin.Context) {
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

// "Get By Id" godoc
// @Summary Get Tag By Id
// @Description Fetch one Tag By Id
// @Tags tag
// @Produce json
// @Param id path integer true "Id of the tag to fetch"
// @Success 200 {object} controllers.Response{data=models.Tag}
// @Router /v1/tag/{id} [get]
func (t *tagController) FindById(c *gin.Context) {
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

// "Create" godoc
// @Summary Create tag
// @Description Create tag
// @Tags tag
// @Accept json
// @Produce json
// @Param createDto body controllers.CreateTagDto true "Body of the tag to create"
// @Success 200 {object} controllers.Response{data=models.Tag}
// @security ApiKeyAuth
// @Router /v1/tag [post]
func (t *tagController) Create(c *gin.Context) {
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

// "Update" godoc
// @Summary Update tag
// @Description Update tag by Id
// @Tags tag
// @Accept json
// @Produce json
// @Param id path integer true "id of the tag to update"
// @Param updateDto body controllers.UpdateTagDto true "Body of the tag to update"
// @Success 200 {object} controllers.Response{data=controllers.Abr{rowsAffected=int}}
// @security ApiKeyAuth
// @Router /v1/tag/{id} [post]
func (t *tagController) Update(c *gin.Context) {
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

// "Delete" godoc
// @Summary Delete Tag
// @Description Delete Tag by ID
// @Tags tag
// @Produce json
// @Param id path integer true "Id of the tag to delete"
// @Success 200 {object} controllers.Response{data=controllers.Abr{rowAffected=int}}
// @security ApiKeyAuth
// @Router /v1/tag/{id} [delete]
func (t *tagController) Delete(c *gin.Context) {
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

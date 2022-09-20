package controllers

import (
	"net/http"
	"strconv"

	"github.com/DustinDust/gin-blog-post/models"
	"github.com/gin-gonic/gin"
)

type blogPostController struct{}

type CreateBlogPostDto struct {
	Title   string       `json:"title" binding:"required"`
	Content string       `json:"content" binding:"required"`
	Tags    []models.Tag `json:"tags" binding:"required"`
}

type UpdateBlogPostDto struct {
	Title   string       `json:"title,omitempty"`
	Content string       `json:"content,omitempty"`
	Tags    []models.Tag `json:"tags,omitempty"`
}

func (b *blogPostController) FindAll(c *gin.Context) {
	query := make(map[string]interface{})
	urlQueries := c.Request.URL.Query()

	// god bless go because i have to actually write this. When will the generic be adapted to stdlib :(
	for key, elem := range urlQueries {
		query[key] = elem
	}

	res, blogPosts := models.BlogPostRepository.FindAll(query)
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
		Data:    blogPosts,
	})
}

func (b *blogPostController) FindById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res, blogPost := models.BlogPostRepository.FindById(id)
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
		Data:    blogPost,
	})
}

func (b *blogPostController) Create(c *gin.Context) {
	createDto := CreateBlogPostDto{}
	if err := c.ShouldBindJSON(&createDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	blogPost := models.BlogPost{
		Title:   createDto.Title,
		Content: createDto.Content,
		Tags:    createDto.Tags,
	}

	res, blogPost := models.BlogPostRepository.Create(blogPost)
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
		Data:    blogPost,
	})
}

func (b *blogPostController) Update(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	updateDto := UpdateBlogPostDto{}
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	// log.Printf("%v", updateDto)
	blogPost := models.BlogPost{
		Content: updateDto.Content,
		Title:   updateDto.Title,
		Tags:    updateDto.Tags,
	}
	res := models.BlogPostRepository.Update(id, blogPost)
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

func (b *blogPostController) Delete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	res := models.BlogPostRepository.Delete(id)
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

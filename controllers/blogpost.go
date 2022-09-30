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

// "Get All" godoc
// @Summary Get All BlogPost
// @Description Fetch all Blogpost based on query and paging
// @Tags blog-post
// @Produce json
// @Param page query integer true "Which page to get"
// @Success 200 {object} controllers.Response{data=[]models.Pagination[models.BlogPost]}
// @Router /v1/blog-post [get]
func (b *blogPostController) FindAll(c *gin.Context) {
	query := make(map[string]interface{})
	urlQueries := c.Request.URL.Query()

	// god bless go because I have to actually write this. When will the generic be adapted to stdlib :(
	for key, elem := range urlQueries {
		query[key] = elem
	}

	res, data := models.BlogPostRepository.FindAll(query)
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
		Data:    data,
	})
}

// "Get By Id" godoc
// @Summary Get BlogPost By Id
// @Description Fetch Blogpost based on ID
// @Tags blog-post
// @Produce json
// @Param id path integer true "Id of the blogpost to fetch"
// @Success 200 {object} controllers.Response{data=models.BlogPost}
// @Router /v1/blog-post/{id} [get]
func (b *blogPostController) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
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

// "Create" godoc
// @Summary Create Blogpost
// @Description Create Blogpost
// @Tags blog-post
// @Accept json
// @Produce json
// @Param createDto body controllers.CreateBlogPostDto true "Body of the blogpost to create"
// @Success 200 {object} controllers.Response{data=models.BlogPost}
// @Security  ApiKeyAuth
// @Router /v1/blog-post [post]
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

// "Update" godoc
// @Summary Update Blogpost
// @Description Update Blogpost
// @Tags blog-post
// @Accept json
// @Produce json
// @Param updateDto body controllers.UpdateBlogPostDto true "Body of the blogpost to update"
// @Param id path integer true "Id of the blogpost to update"
// @Success 200 {object} controllers.Response{data=controllers.Abr{rowsAffected=int}}
// @Router /v1/blog-post/{id} [put]
// @Security ApiKeyAuth
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

// "Delete" godoc
// @Summary Delete Blogpost
// @Description Delete Blogpost by ID
// @Tags blog-post
// @Produce json
// @Param id path integer true "Id of the blogpost to delete"
// @Success 200 {object} controllers.Response{data=controllers.Abr{rowAffected=int}}
// @security ApiKeyAuth
// @Router /v1/blog-post/{id} [delete]
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

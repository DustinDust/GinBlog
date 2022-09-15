package controllers

var UserController *UserControllerInterface
var BlogPostController *BlogPostControllerInterface
var TagController *TagControllerInterface

type Response struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func InitController() {
	UserController = &UserControllerInterface{}
	BlogPostController = &BlogPostControllerInterface{}
	TagController = &TagControllerInterface{}
}

package controllers

var UserController *userController
var BlogPostController *blogPostController
var TagController *tagController
var AuthController *authController

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Abr struct {
}

func InitController() {
	UserController = &userController{}
	BlogPostController = &blogPostController{}
	TagController = &tagController{}
	AuthController = &authController{}
}

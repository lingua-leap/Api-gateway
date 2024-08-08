package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(c *gin.Context)
	GetUserProfile(c *gin.Context)
	GetAllUsers(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
	ChangePassword(c *gin.Context)
	DeleteUser(c *gin.Context)
}

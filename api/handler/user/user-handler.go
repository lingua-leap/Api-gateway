package user

import (
	pb "api-gateway/generated/user"
	"api-gateway/pkg/token"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type User struct {
	user pb.UserServiceClient
	log  *slog.Logger
}

func NewUser(user pb.UserServiceClient, log *slog.Logger) *User {
	return &User{user: user, log: log}
}

func (u *User) CreateUser(c *gin.Context) {
	var user *pb.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	res, err := u.user.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *User) GetUserProfile(c *gin.Context) {
	var req *pb.UserId

	value, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		u.log.Error("unauthorized claims")
		return
	}

	claims, a := value.(*token.Claims)
	if !a {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		u.log.Error("unauthorized claims")
		return
	}

	req.Id = claims.Id
	res, err := u.user.GetUserProfile(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *User) GetAllUsers(c *gin.Context) {
	var req *pb.FilterRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	res, err := u.user.GetAllUsers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *User) UpdateUserProfile(c *gin.Context) {
	var req *pb.UpdateUserPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	value, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		u.log.Error("unauthorized claims")
		return
	}

	claims, a := value.(*token.Claims)
	if !a {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		u.log.Error("unauthorized claims")
		return
	}

	req.Id = claims.Id

	res, err := u.user.UpdateUserProfile(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *User) ChangePassword(c *gin.Context) {
	var req *pb.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return

	}

	res, err := u.user.ChangePassword(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *User) DeleteUser(c *gin.Context) {
	var req *pb.UserId
	value, ok := c.Get("claims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		u.log.Error("unauthorized claims")
		return
	}

	claims, a := value.(*token.Claims)
	if !a {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		u.log.Error("unauthorized claims")
		return
	}

	req.Id = claims.Id
	res, err := u.user.DeleteUser(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		u.log.Error("error in create user", "error", err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

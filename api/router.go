package api

import (
	"api-gateway/api/handler"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Init()
	Run() error
}

func NewRouter(hd handler.Handler) Router {
	router := gin.Default()
	return &RouterImpl{handler: hd, router: router}
}

type RouterImpl struct {
	handler handler.Handler
	router  *gin.Engine
}

func (r *RouterImpl) Init() {
	router := r.router.Group("api")

	us := r.handler.NewUserHandler()
	user1 := router.Group("user")
	user1.POST("/create", us.CreateUser) //admin
	user1.GET("/all", us.GetAllUsers)    //admin
	user1.GET("", us.GetUserProfile)
	user1.PUT("", us.UpdateUserProfile)
	user1.PUT("/change-password", us.ChangePassword)
	user1.DELETE("", us.DeleteUser)

}

func (r *RouterImpl) Run() error {
	return r.router.Run(":8080")
}

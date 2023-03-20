package router

import (
	handler "go-fp-crowdfundchat/delivery/http"
	"go-fp-crowdfundchat/delivery/middleware"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(u *handler.UserHandler, p *handler.ProjectHandler) {
	r = gin.Default()
	r.Static("/images", "./mock/images")
	api := r.Group("v1/api")

	api.POST("/user/register", u.RegisterUser)
	api.POST("/user/login", u.LoginUser)
	
	apiAuth := r.Group("v1/api")
	api.POST("/user/pin-validation", u.PinValidation)
	apiAuth.Use(middleware.AuthMiddleWare())
	apiAuth.POST("/user/verify-phone", u.IsPhoneNoAvailable)
	apiAuth.POST("/user/upload-image", u.UploadImage)
	apiAuth.GET("/user/verify-user", u.VerifyUser)

	apiAuth.GET("/project/project-list", p.ProjectList)
	apiAuth.GET("/project/project-detail/:id", p.ProjectDetail)
	/*r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))*/

	/*r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)*/
}

func Start() error {
	return r.Run(":8080")
}
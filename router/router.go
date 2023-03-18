package router

import (
	handler "go-fp-crowdfundchat/delivery/http"
	"go-fp-crowdfundchat/delivery/middleware"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(handler *handler.UserHandler) {
	r = gin.Default()
	api := r.Group("v1/api")

	api.POST("/user/register", handler.RegisterUser)
	api.POST("/user/login", handler.LoginUser)
	
	apiAuth := r.Group("v1/api")
	//api.POST("/user/verify-phone", handler.IsPhoneNoAvailable)
	api.POST("/user/pin-validation", handler.PinValidation)
	apiAuth.Use(middleware.AuthMiddleWare())
	apiAuth.POST("/user/verify-phone", handler.IsPhoneNoAvailable)
	apiAuth.POST("/user/upload-image", handler.UploadImage)

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
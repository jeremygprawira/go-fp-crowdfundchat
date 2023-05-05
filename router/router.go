package router

import (
	handler "go-fp-crowdfundchat/delivery/http"
	"go-fp-crowdfundchat/delivery/middleware"
	"go-fp-crowdfundchat/delivery/websockets"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(u *handler.UserHandler, p *handler.ProjectHandler, t *handler.TransactionHandler, c *handler.ChatHandler, w *websockets.WebsocketHandler) {
	//go handler.Broadcaster()
	r = gin.Default()
	r.Use(cors.Default())
	r.Static("/images", "./mock/images")
	api := r.Group("v1/api")

	api.POST("/user/register", u.RegisterUser)
	api.POST("/user/login", u.LoginUser)
	api.POST("/user/pin-validation", u.PinValidation)
	
	apiAuth := r.Group("v1/api")
	apiAuth.Use(middleware.AuthMiddleWare())

	apiAuth.POST("/user/verify-phone", u.IsPhoneNoAvailable)
	apiAuth.POST("/user/upload-image", u.UploadImage)
	apiAuth.GET("/user/verify-user", u.VerifyUser)

	apiAuth.GET("/project/project-list", p.ProjectList)
	apiAuth.GET("/project/project-detail/:id", p.ProjectDetail)
	apiAuth.POST("/project/create-project", p.CreateProject)
	apiAuth.PUT("/project/update-project/:id", p.UpdateProject)
	apiAuth.POST("/project/upload-project-image", p.UploadProjectImage)

	apiAuth.GET("/project/:id/transaction", t.ProjectTransactionList)
	apiAuth.GET("/user/transaction", t.UserTransactionList)
	apiAuth.POST("/transaction/order", t.OrderTransaction)
	api.POST("/transaction/verify", t.VerifyTransaction)

	api.POST("/chat/send", c.InputChat)
	api.GET("/chat/history/:origin_id/:destination_id", c.ChatHistory)
	api.GET("/chat/:user_id", handler.InitiateChat)
	//api.GET("/ws", handler.HandleConnection)

	//api.POST("/ws/create-room", w.CreateRoom)
}

func Start() error {
	return r.Run(":8080")
}
/*
func RunWebsocket(c *handler.ChatHandler) error {
	go handler.Broadcaster()
	
	r = gin.Default()
	r.Use(cors.Default())
	api := r.Group("v1/api")
	api.GET("/ws", handler.HandleConnection)

	return r.Run(":8081")
}*/
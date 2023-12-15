package router

import (
	handler "go-fp-crowdfundchat/delivery/http"
	"go-fp-crowdfundchat/delivery/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(u *handler.UserHandler, p *handler.ProjectHandler, t *handler.TransactionHandler, c *handler.ChatHandler) {
	r = gin.Default()
	r.Use(cors.Default())
	r.Static("/images", "./mock/images")
	api := r.Group("v1/api")

	api.POST("/user/register", u.RegisterUser)
	api.POST("/user/login", u.LoginUser)
	api.POST("/user/pin-validation", u.PinValidation)
	api.PUT("/user/update", u.UpdateUser)
	
	apiAuth := r.Group("v1/api")
	apiAuth.Use(middleware.AuthMiddleWare())
	
	apiAuth.POST("/user/verify-phone", u.IsPhoneNoAvailable)
	apiAuth.POST("/user/upload-image", u.UploadImage)
	apiAuth.GET("/user/verify-user", u.VerifyUser)

	api.GET("/project/project-list", p.ProjectList)
	apiAuth.GET("/project/project-detail/:id", p.ProjectDetail)
	apiAuth.POST("/project/create-project", p.CreateProject)
	apiAuth.PUT("/project/update-project/:id", p.UpdateProject)
	apiAuth.POST("/project/upload-project-image", p.UploadProjectImage)

	apiAuth.GET("/project/:id/transaction", t.ProjectTransactionList)
	apiAuth.GET("/user/transaction", t.UserTransactionList)
	apiAuth.POST("/transaction/order", t.OrderTransaction)
	api.POST("/transaction/verify", t.VerifyTransaction)

	api.POST("/chat/create", c.CreateChatRoom)
	apiAuth.GET("/chat/join/:id", c.WebsocketJoinRoom)
	api.GET("/chat/rooms", c.GetRoom)
	api.GET("/chat/clients/:id", c.GetClient)
	api.GET("/chat/chat/:id", c.ChatRoom)
	api.GET("/chat/history/:id", c.ChatHistory)
}

func Start() error {
	return r.Run(":8080")
}
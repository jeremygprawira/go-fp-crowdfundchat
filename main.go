package main

import (
	"go-fp-crowdfundchat/database"
	handler "go-fp-crowdfundchat/delivery/http"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/router"
	"go-fp-crowdfundchat/usecase"
	"log"
)

func main() {
    dbConnection, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("could not initialize database connection: %s", err)
    }
    
    baseRepository := repository.NewBaseRepository(dbConnection)
	
    userUsecase := usecase.NewUserUsecase(baseRepository)
    userHandler := handler.NewUserHandler(userUsecase)
	
    projectUsecase := usecase.NewProjectUsecase(baseRepository)
    projectHandler := handler.NewProjectHandler(projectUsecase, userUsecase)

    transactionUsecase := usecase.NewTransactionUsecase(baseRepository)
    transactionHandler := handler.NewTransactionHandler(transactionUsecase, userUsecase)

    chatUsecase := usecase.NewChatUsecase(baseRepository)
    chatHub := handler.NewHub(chatUsecase)
    chatHandler := handler.NewChatHandler(chatHub, userUsecase, projectUsecase, chatUsecase)
    go chatHub.Run()
    
    router.InitRouter(userHandler, projectHandler, transactionHandler, chatHandler)
    router.Start()

}

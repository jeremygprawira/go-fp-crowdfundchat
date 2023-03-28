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

	//userRepository := repository.NewBaseRepository(dbConnection)
	userUsecase := usecase.NewUserUsecase(baseRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	//projectRepository := repository.NewBaseRepository(dbConnection)
	projectUsecase := usecase.NewProjectUsecase(baseRepository)
	projectHandler := handler.NewProjectHandler(projectUsecase, userUsecase)

	transactionUsecase := usecase.NewTransactionUsecase(baseRepository)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase, userUsecase)

	router.InitRouter(userHandler, projectHandler, transactionHandler)
	router.Start()

}
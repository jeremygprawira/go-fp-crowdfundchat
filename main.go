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
	
	userRepository := repository.NewBaseRepository(dbConnection)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	projectRepository := repository.NewBaseRepository(dbConnection)
	projectUsecase := usecase.NewProjectUsecase(projectRepository)
	projectHandler := handler.NewProjectHandler(projectUsecase, userUsecase)

	router.InitRouter(userHandler, projectHandler)
	router.Start()

}
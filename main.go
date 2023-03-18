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
	
	userRepository := repository.NewUserRepository(dbConnection)
	userUsecase := usecase.NewUserUsecase(userRepository)

	userHandler := handler.NewUserHandler(userUsecase)

	router.InitRouter(userHandler)
	router.Start()

}
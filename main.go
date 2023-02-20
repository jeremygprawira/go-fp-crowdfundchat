package main

import (
	"go-fp-crowdfundchat/database"
	"go-fp-crowdfundchat/handler"
	"go-fp-crowdfundchat/router"
	"go-fp-crowdfundchat/usecase/user"
	"log"
)

func main() {
	dbConnection, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRepository := user.NewRepository(dbConnection)
	userService := user.NewService(userRepository)

	/*userInput := user.RegisterUserRequest{}
	userInput.Name = "TEST-NEW-4"
	userInput.PIN = "TEST-NEW-4"
	userInput.PhoneNo = "TEST-NEW-4"

	userService.PostRegisterUser(userInput)*/

	userHandler := handler.NewUserHandler(userService)
	router.InitRouter(userHandler)
	router.Start()

}
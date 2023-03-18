package main

import (
	"fmt"
	"go-fp-crowdfundchat/database"
	handler "go-fp-crowdfundchat/delivery/http"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/router"
	"go-fp-crowdfundchat/usecase"
	"go-fp-crowdfundchat/util"
	"log"
)

func main() {
	dbConnection, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	
	userRepository := repository.NewUserRepository(dbConnection)
	userUsecase := usecase.NewService(userRepository)
	
	a := "222232"
	b, err := util.RepeatingPINChecker(a)
	if err != nil {
		fmt.Println("nt")
		return
	}
	fmt.Println(b)

	userHandler := handler.NewUserHandler(userUsecase)
	/*userInput := user.RegisterUserRequest{}
	userInput.Name = "TEST-NEW-4"
	userInput.PIN = "TEST-NEW-4"
	userInput.PhoneNo = "TEST-NEW-4"

	userService.PostRegisterUser(userInput)*/

	router.InitRouter(userHandler)
	router.Start()

}
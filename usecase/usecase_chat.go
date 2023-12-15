package usecase

import (
	"encoding/base64"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
	"strconv"
)

type ChatUsecase interface{
	PostInputChat(request *model.ChatRequest) (*model.Chat, error)
	GetChatByProjectID(projectID int) ([]*model.Chat, error)
	//GetChatHistory(request *model.ChatHistoryRequest) ([]*model.Chat, error)
}

type chatUsecase struct {
	repo repository.BaseRepository
}

func NewChatUsecase(repo repository.BaseRepository) *chatUsecase {
	return &chatUsecase{repo}
}

func (ch *chatUsecase) PostInputChat(request *model.ChatRequest) (*model.Chat, error) {
	chat := model.Chat{}
	userID, _ := strconv.Atoi(request.UserID)
	chat.UserID = userID
	
	user, err := ch.repo.FindUserByID(userID) 
	if err != nil {
		return nil, err
	}

	chat.UserName = user.Name

	chat.ProjectID, _ = strconv.Atoi(request.ProjectID)
	//encryptedChat, err := util.HashEncrypt([]byte(request.Message))
	encryptedChat := util.Encrypt([]byte(request.Chat))
	/*if err != nil {
		return &chat, err
	}*/
	chat.Content = base64.StdEncoding.EncodeToString(encryptedChat)

	newChat, err := ch.repo.CreateChatToDB(&chat)
	if err != nil {
		return newChat, err
	}

	return newChat, nil
}

func (ch *chatUsecase) GetChatByProjectID(projectID int) ([]*model.Chat, error) {
	chat, err := ch.repo.FindChatByProjectID(projectID)
	if err != nil {
		return chat, nil
	}

	return chat, nil
}
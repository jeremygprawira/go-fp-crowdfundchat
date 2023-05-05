package usecase

import (
	"encoding/base64"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/repository"
	"go-fp-crowdfundchat/util"
)

type ChatUsecase interface{
	PostInputChat(request *model.InputChatRequest) (*model.Chat, error)
	GetChatHistory(request *model.ChatHistoryRequest) ([]*model.Chat, error)
}

type chatUsecase struct {
	repo repository.BaseRepository
}

func NewChatUsecase(repo repository.BaseRepository) *chatUsecase {
	return &chatUsecase{repo}
}

func (ch *chatUsecase) PostInputChat(request *model.InputChatRequest) (*model.Chat, error) {
	chat := model.Chat{}
	chat.OriginUserID = request.OriginUserID
	chat.DestinationUserID = request.DestinationUserID
	//encryptedChat, err := util.HashEncrypt([]byte(request.Message))
	encryptedChat := util.Encrypt([]byte(request.Message))
	/*if err != nil {
		return &chat, err
	}*/
	chat.Message = base64.StdEncoding.EncodeToString(encryptedChat)

	newChat, err := ch.repo.CreateChatToDB(&chat)
	if err != nil {
		return newChat, err
	}

	return newChat, nil
}

func (ch *chatUsecase) GetChatHistory(request *model.ChatHistoryRequest) ([]*model.Chat, error) {
	chat := model.InputChatRequest{}
	chat.OriginUserID = request.OriginUserID
	chat.DestinationUserID = request.DestinationUserID
	newHistory, err := ch.repo.FindChatByIDs(chat.OriginUserID, chat.DestinationUserID)
	//chat, err := ch.repo.FindChatByIDs(request.OriginUserID, request.DestinationUserID)
	if err != nil {
		return newHistory, err
	}
	return newHistory, nil
}

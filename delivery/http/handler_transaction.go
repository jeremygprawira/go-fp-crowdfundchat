package http

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase	usecase.TransactionUsecase	
}

func NewTransactionHandler (transactionUsecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUsecase}
}

func (h *TransactionHandler) ProjectTransactionList(c *gin.Context) {
	var request *model.ProjectTransactionListRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40019",
			"responseMessage": "Failed to retrieve the project's transaction",
		})
		return
	}

	transaction, err := h.transactionUsecase.GetTransactionByProjectID(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40020",
			"responseMessage": "Failed to retrieve the project's transaction",
		})
		return
	}

	transactionList := make([]*model.ProjectTransactionListResponse, len(transaction))
	for i, p := range transaction {
		transactionList[i] = &model.ProjectTransactionListResponse{
			ID: p.ID,
			UserName: p.User.Name,
			Amount: p.Amount,
			CreatedAt: p.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"transactions": transactionList,
	})
}

func (h *TransactionHandler) UserTransactionList(c *gin.Context) {
	userID := c.GetInt("currentUser")
	
	transaction, err := h.transactionUsecase.GetTransactionByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40021",
			"responseMessage": "Failed to retrieve the user's transaction",
		})
		return
	}

	var projectImagesList *model.ProjectImagesTransactionListResponse
	for _, p := range transaction {
		if len(p.Project.ProjectImages) > 0 {
			projectImagesList = &model.ProjectImagesTransactionListResponse{
				Name: p.Project.ProjectTitle,
				ImageURL: p.Project.ProjectImages[0].FileName,
			}
		} else if len(p.Project.ProjectImages) == 0 {
			projectImagesList = &model.ProjectImagesTransactionListResponse{
				Name: p.Project.ProjectTitle,
				ImageURL: "",
			}
		}
	}

	userTransactionList := make([]*model.UserTransactionListResponse, len(transaction))
	for i, p := range transaction {
		userTransactionList[i] = &model.UserTransactionListResponse{
			ID: p.ID,
			Amount: p.Amount,
			Status: p.Status,
			CreatedAt: p.CreatedAt,
			Project: *projectImagesList,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"transactions": userTransactionList,
	})
}
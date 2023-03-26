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
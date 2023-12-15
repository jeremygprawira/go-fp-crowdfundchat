package http

import (
	"fmt"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
    transactionUsecase  usecase.TransactionUsecase
    userUsecase usecase.UserUsecase 
}

func NewTransactionHandler (transactionUsecase usecase.TransactionUsecase, userUsecase usecase.UserUsecase) *TransactionHandler {
    return &TransactionHandler{transactionUsecase, userUsecase}
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

    request.User.ID = c.GetInt("currentUser")

    transaction, err := h.transactionUsecase.GetTransactionByProjectID(request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "responseCode": "40020",
            "responseMessage": err.Error(),
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
        "data": transactionList,
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
        "data": userTransactionList,
    })
}

func (h *TransactionHandler) OrderTransaction(c *gin.Context) {
    var request model.OrderTransactionRequest

    err := c.ShouldBindJSON(&request)
    if err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "responseCode": "42209",
            "responseMessage": "The required field on the body request is empty or invalid.",
        })
        return
    }

    userID := c.GetInt("currentUser")
    user, err := h.userUsecase.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "responseCode": "42210",
            "responseMessage": "The required field on the body request is empty or invalid.",
        })
        return
    }

    request.User = *user

    transaction, err := h.transactionUsecase.PostOrderTransaction(&request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "responseCode": "40022",
            "responseMessage": fmt.Sprintf("Usecase PostOrderTransaction is not working properly: %s", err.Error()),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "responseCode": "20000",
        "responseMessage": "Request has been successfully sent.",
        "transaction_id": transaction.ID,
        "project_id": transaction.ProjectID,
        "user_id": transaction.UserID,
        "amount": transaction.Amount,
        "status": transaction.Status,
        "receipt_no": transaction.ReceiptNo,
        "transaction_url": transaction.TransactionURL,
    })
}

func (h *TransactionHandler) VerifyTransaction(c *gin.Context) {
	var request model.TransactionNotificationRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "responseCode": "40023",
            "responseMessage": fmt.Sprintf("Failed to process the transaction: %s", err.Error()),
        })
        return
    }

	err = h.transactionUsecase.PostVerifyTransaction(&request)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "responseCode": "40024",
            "responseMessage": fmt.Sprintf("Failed to process the transaction: %s", err.Error()),
        })
        return
    }

	c.JSON(http.StatusOK, gin.H{
        "responseCode": "20000",
        "responseMessage": "Request has been successfully sent.",
    })
}
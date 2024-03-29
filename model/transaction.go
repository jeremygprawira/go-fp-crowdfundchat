package model

import "time"

type Transaction struct {
    ID                      int
    ProjectID               int                                         `json:"project_id"`
    UserID                  int                                         `json:"user_id"`
    Amount                  int                                         `json:"amount"`
    Status                  string                                      `json:"status"`
    ReceiptNo               string                                      `json:"receipt_no"`
    TransactionURL          string                                      `json:"transaction_url"`
    CreatedAt               time.Time                                   `json:"created_at"`
    UpdatedAt               time.Time                                   `json:"updated_at"`
    User                    User                                        `gorm:"foreignKey:UserID"`
    Project                 Project
}

type ProjectTransactionListRequest struct {
    ID                      int                                         `uri:"id" binding:"required"`
    User                    User
}

type ProjectTransactionListResponse struct {
    ID                      int
    UserName                string                                      `json:"user_name"`
    Amount                  int                                         `json:"amount"`
    CreatedAt               time.Time                                   `json:"created_at"`
}

type UserTransactionListResponse struct {
    ID                      int
    Amount                  int                                         `json:"amount"`
    Status                  string                                      `status:"status"`
    CreatedAt               time.Time                                   `json:"created_at"`
    Project                 ProjectImagesTransactionListResponse        `json:"project"`
}

type ProjectImagesTransactionListResponse struct {
    Name                    string          							`json:"name"`
    ImageURL                string          							`json:"image_url"`
}

type OrderTransactionRequest struct {
    ProjectID               int                                         `json:"project_id" binding:"required"`
    Amount                  int                                         `json:"amount" binding:"required"`
    User                    User                                        `gorm:"foreignKey:UserID"`
}

type TransactionNotificationRequest struct {
	TransactionStatus		string										`json:"transaction_status"`
	OrderID					string										`json:"order_id"`
	PaymentType				string										`json:"payment_type"`
	FraudStatus				string										`json:"fraud_status"`
}
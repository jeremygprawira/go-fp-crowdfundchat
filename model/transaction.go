package model

import "time"

type Transaction struct {
	ID						int
	ProjectID				int				`json:"project_id"`
	UserID					int				`json:"user_id"`
	Amount					int				`json:"amount"`
	Status					string			`json:"status"`
	Code					string			`json:"code"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
	User					User			`gorm:"foreignKey:UserID"`
}

type ProjectTransactionListRequest struct {
	ID						int				`uri:"id" binding:"required"`
	User					User
}

type ProjectTransactionListResponse struct {
	ID						int
	UserName				string			`json:"user_name"`
	Amount					int				`json:"amount"`
	CreatedAt 				time.Time		`json:"created_at"`
}
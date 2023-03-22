package model

import "time"

type Transaction struct {
	ID						int
	ProjectID				int
	UserID					int
	Amount					int
	Status					string
	Code					string
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
}
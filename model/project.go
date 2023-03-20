package model

import (
	"time"
)

type Project struct {
	ID 						int
	UserID					int				`json:"user_id"`	
	ProjectTitle 			string 			`json:"project_title"`
	ShortDescription 		string			`json:"short_description"`
	LongDescription			string			`json:"long_description"`
	Perks 					string			`json:"perks"`
	Slug					string			`json:"slug"`
	ContributorCount 		int				`json:"contributor_count"`
	GoalAmount 				int				`json:"goal_amount"`
	CurrentAmount			int				`json:"current_amount"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`

	ProjectImages			[]ProjectImages	`json:"project_images"`
}

type ProjectImages struct {
	ID 						int				
	ProjectID 				int				`json:"project_id"`
	FileName 				string			`json:"project_image_file_name"`
	IsPrimary				int				`json:"is_primary"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
}

type ProjectListResponse struct {
	ID 						int				`json:"id"`	
	UserID					int				`json:"user_id"`	
	ProjectTitle 			string 			`json:"project_title"`
	ShortDescription 		string			`json:"short_description"`
	Slug					string			`json:"slug"`
	ContributorCount 		int				`json:"contributor_count"`
	GoalAmount 				int				`json:"goal_amount"`
	CurrentAmount			int				`json:"current_amount"`
	
	ProjectImages			[]ProjectImages	`json:"project_images" gorm:"foreignKey:ProjectID"`
}

type ProjectDetailRequest struct {
	ID						int				`uri:"id" binding:"required"`
}

type ProjectDetailResponse struct {
	ID 						int				`json:"id"`	
	UserID					int				`json:"user_id"`	
	ProjectTitle 			string 			`json:"project_title"`
	ShortDescription 		string			`json:"short_description"`
	LongDescription			string			`json:"long_description"`
	Slug					string			`json:"slug"`
	ContributorCount 		int				`json:"contributor_count"`
	GoalAmount 				int				`json:"goal_amount"`
	CurrentAmount			int				`json:"current_amount"`
	Perks 					string			`json:"perks"`

	User					[]User			`json:"user" gorm:"foreignKey:ID"`
	ProjectImages			[]ProjectImages	`json:"project_images" gorm:"foreignKey:ProjectID"`
}

/*type ProjectDetailResponse struct {
	ID 						int				`json:"id"`	
	UserID					int				`json:"user_id"`	
	ProjectTitle 			string 			`json:"project_title"`
	ShortDescription 		string			`json:"short_description"`
	LongDescription			string			`json:"long_description"`
	Slug					string			`json:"slug"`
	ContributorCount 		int				`json:"contributor_count"`
	GoalAmount 				int				`json:"goal_amount"`
	CurrentAmount			int				`json:"current_amount"`
	Perks 					string			`json:"perks"`

	User					[]User 
	ProjectImages			[]ProjectImages	`json:"project_images" gorm:"foreignKey:ProjectID"`
}*/

type CreateProjectInput struct {
	ProjectTitle 			string 			`json:"project_title"`
	ShortDescription 		string			`json:"short_description"`
	LongDescription			string			`json:"long_description"`
	GoalAmount 				int				`json:"goal_amount"`
	Perks 					string			`json:"perks"`
	User					User
}
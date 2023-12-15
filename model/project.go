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
	User					User			`json:"user" gorm:"foreignKey:ID"`

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
	ImageURL				string			`json:"image_url"`
	
	//ProjectImages			[]ProjectImages	`json:"project_images" gorm:"foreignKey:ProjectID"`
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

type ProjectImagesProjectDetailResponse struct {
    IsPrimary               int         	`json:"is_primary"`
    ImageURL                string          `json:"image_url"`
}

type CreateProjectRequest struct {
	ProjectTitle 			string 			`json:"project_title" binding:"required"`
	ShortDescription 		string			`json:"short_description" binding:"required"`
	LongDescription			string			`json:"long_description" binding:"required"`
	GoalAmount 				int				`json:"goal_amount" binding:"required"`
	Perks 					string			`json:"perks" binding:"required"`
	User					User			`json:"user" gorm:"foreignKey:ID"`
}

type UpdateProjectRequest struct {
	ID						int				`uri:"id" binding:"required"`
}

type UploadProjectImageRequest struct {
	ProjectID				int				`form:"project_id" binding:"required"`
	IsPrimary				bool			`form:"is_primary"`
	User					User			`json:"user" gorm:"foreignKey:ID"`
}
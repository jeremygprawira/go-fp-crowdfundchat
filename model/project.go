package model

import "time"

type Project struct {
	ID 						int
	UserID					int				`json:"user_id"`	
	ProjectTitle 			string 			`json:"project_title"`
	ShortDescription 		string			`json:"short_description"`
	LongDescription			string			`json:"long_description"`
	Perks 					string			`json:"perks"`
	ContributorCount 		int				`json:"contributor_count"`
	GoalAmount 				int				`json:"goal_amount"`
	CurrentAmount			int				`json:"current_amount"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
}

type ProjectImage struct {
	ID 						int
	ProjectID 				int				`json:"project_id"`
	ProjectImageFileName 	string			`json:"project_image_file_name"`
	isPrimary				int				`json:"is_primary"`
	CreatedAt 				time.Time		`json:"created_at"`
	UpdatedAt 				time.Time		`json:"updated_at"`
}
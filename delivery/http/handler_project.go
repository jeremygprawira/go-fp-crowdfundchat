package http

import (
	"fmt"
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectUsecase	usecase.ProjectUsecase
	userUsecase usecase.UserUsecase
}

func NewProjectHandler(projectUsecase usecase.ProjectUsecase, userUsecase usecase.UserUsecase) *ProjectHandler {
	return &ProjectHandler{projectUsecase, userUsecase}
}

func (h *ProjectHandler) ProjectList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	projects, err := h.projectUsecase.GetProjectList(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40009",
			"responseMessage": "Failed to get project list",
		})
		return
	}
	projectList := make([]*model.ProjectListResponse, len(projects))
	for i, p := range projects {
		
		if len(p.ProjectImages) > 0 {
			projectList[i] = &model.ProjectListResponse{
				ID: p.ID,
				ProjectTitle: p.ProjectTitle,
				UserID: p.UserID,
				ShortDescription: p.ShortDescription,
				ImageURL: p.ProjectImages[0].FileName,
				GoalAmount: p.GoalAmount,
				CurrentAmount: p.CurrentAmount,
			}
		} else if len(p.ProjectImages) == 0 {
			projectList[i] = &model.ProjectListResponse{
				ID: p.ID,
				ProjectTitle: p.ProjectTitle,
				UserID: p.UserID,
				ShortDescription: p.ShortDescription,
				ImageURL: "",
				GoalAmount: p.GoalAmount,
				CurrentAmount: p.CurrentAmount,
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Project list has been successfully retrieved.",
		"data": projectList,
	})
}

func (h *ProjectHandler) ProjectDetail(c *gin.Context) {
	var request *model.ProjectDetailRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40010",
			"responseMessage": "Failed to get project detail",
		})
		return
	}

	project, err := h.projectUsecase.GetProjectDetail(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40011",
			"responseMessage": "Failed to get project detail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Project detail has been successfully retrieved.",
		"data": &project,
	})
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var projectRequest model.CreateProjectRequest

	err := c.ShouldBindJSON(&projectRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42205",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	userID := c.GetInt("currentUser")
	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42206",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	projectRequest.User = *user

	project, err := h.projectUsecase.PostCreateProject(&projectRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40012",
			"responseMessage": "Usecase PostCreateProject is not working properly",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Project has been created successfully",
		"project_id": project.ID,
		"user_id": project.UserID,
		"project_title": project.ProjectTitle,
		"short_description": project.ShortDescription,
		"slug": project.Slug,
		"current_amount": project.CurrentAmount,
		"goal_amount": project.GoalAmount,
		"image_url": project.ProjectImages,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var request *model.UpdateProjectRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40013",
			"responseMessage": "ID is not found",
		})
		return
	}

	var requestData *model.CreateProjectRequest
	err = c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42207",
			"responseMessage": "Failed to update project",
		})
		return
	}

	userID := c.GetInt("currentUser")
	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42208",
			"responseMessage": "Could not get the User ID",
		})
		return
	}

	requestData.User = *user 

	updatedCampaign, err := h.projectUsecase.PutUpdateProject(request, requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40014",
			"responseMessage": "Failed to update project",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Project has been created successfully",
		"project_id": updatedCampaign.ID,
		"user_id": updatedCampaign.UserID,
		"project_title": updatedCampaign.ProjectTitle,
		"short_description": updatedCampaign.ShortDescription,
		"slug": updatedCampaign.Slug,
		"current_amount": updatedCampaign.CurrentAmount,
		"goal_amount": updatedCampaign.GoalAmount,
		"image_url": updatedCampaign.ProjectImages,
	})
}

func (h *ProjectHandler) UploadProjectImage (c *gin.Context) {
	var request model.UploadProjectImageRequest
	
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40015",
			"responseMessage": err.Error(),
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40016",
			"responseMessage": "Failed to upload image to your project",
			"is_uploaded": false,
		})
		return
	}

	userID := c.GetInt("currentUser")

	path := fmt.Sprintf("mock/images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "400017",
			"responseMessage": "Failed to upload image to your project",
		})
		return
	}
	
	_, err = h.projectUsecase.PostUploadProjectImage(&request, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40018",
			"responseMessage": "Failed to upload image to your project",
			"is_uploaded": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully sent.",
		"isImageUploaded": true,
	})
}
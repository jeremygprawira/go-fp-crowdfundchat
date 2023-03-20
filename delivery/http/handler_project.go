package http

import (
	"go-fp-crowdfundchat/model"
	"go-fp-crowdfundchat/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectUsecase	usecase.ProjectUsecase
}

func NewProjectHandler(projectUsecase usecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{projectUsecase}
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

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Project list has been successfully retrieved.",
		"data": projects,
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
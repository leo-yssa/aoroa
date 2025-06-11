package handlers

import (
	"net/http"
	"strconv"

	"aoroa/internal/application"
	"aoroa/pkg/common/errors"

	"github.com/gin-gonic/gin"
)

type IssueHandler struct {
	service *application.IssueService
}

func NewIssueHandler(service *application.IssueService) *IssueHandler {
	return &IssueHandler{service: service}
}

type CreateIssueRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	UserID      *uint  `json:"userId"`
}

type UpdateIssueRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	UserID      *uint   `json:"userId"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (h *IssueHandler) CreateIssue(c *gin.Context) {
	var req CreateIssueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	appReq := application.CreateIssueRequest{
		Title:       req.Title,
		Description: req.Description,
		UserID:      req.UserID,
	}

	issue, err := h.service.CreateIssue(appReq)
	if err != nil {
		status := http.StatusInternalServerError
		if err == errors.ErrUserNotFound {
			status = http.StatusBadRequest
		}
		c.JSON(status, ErrorResponse{
			Error: err.Error(),
			Code:  status,
		})
		return
	}

	c.JSON(http.StatusCreated, issue)
}

func (h *IssueHandler) GetIssues(c *gin.Context) {
	status := c.Query("status")
	issues, err := h.service.GetIssues(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"issues": issues})
}

func (h *IssueHandler) GetIssue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "invalid issue id",
			Code:  http.StatusBadRequest,
		})
		return
	}

	issue, err := h.service.GetIssue(uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err == errors.ErrIssueNotFound {
			status = http.StatusNotFound
		}
		c.JSON(status, ErrorResponse{
			Error: err.Error(),
			Code:  status,
		})
		return
	}

	c.JSON(http.StatusOK, issue)
}

func (h *IssueHandler) UpdateIssue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "invalid issue id",
			Code:  http.StatusBadRequest,
		})
		return
	}

	var req UpdateIssueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	appReq := application.UpdateIssueRequest{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      req.UserID,
	}

	issue, err := h.service.UpdateIssue(uint(id), appReq)
	if err != nil {
		status := http.StatusInternalServerError
		switch err {
		case errors.ErrIssueNotFound:
			status = http.StatusNotFound
		case errors.ErrUserNotFound, errors.ErrInvalidStatus, errors.ErrAssigneeRequired, errors.ErrIssueCompletedOrCancelled:
			status = http.StatusBadRequest
		}
		c.JSON(status, ErrorResponse{
			Error: err.Error(),
			Code:  status,
		})
		return
	}

	c.JSON(http.StatusOK, issue)
} 
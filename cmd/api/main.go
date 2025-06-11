package main

import (
	"aoroa/internal/application"
	"aoroa/internal/infrastructure/repository"
	"aoroa/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize repositories
	issueRepo := repository.NewMemoryIssueRepository()
	userRepo := repository.NewMemoryUserRepository()

	// Initialize services
	issueService := application.NewIssueService(issueRepo, userRepo)

	// Initialize handlers
	issueHandler := handlers.NewIssueHandler(issueService)

	// Initialize router
	r := gin.Default()

	// Register routes
	r.POST("/issue", issueHandler.CreateIssue)
	r.GET("/issues", issueHandler.GetIssues)
	r.GET("/issue/:id", issueHandler.GetIssue)
	r.PATCH("/issue/:id", issueHandler.UpdateIssue)

	// Start server
	r.Run(":8080")
} 
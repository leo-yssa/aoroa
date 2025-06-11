package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Create server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
} 
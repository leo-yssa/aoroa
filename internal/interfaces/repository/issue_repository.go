package repository

import "aoroa/internal/domain"

type IssueRepository interface {
	Create(issue *domain.Issue) error
	FindByID(id uint) (*domain.Issue, error)
	FindAll(status string) ([]*domain.Issue, error)
	Update(issue *domain.Issue) error
}

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	FindAll() ([]*domain.User, error)
} 
package application

import (
	"aoroa/internal/domain"
	"aoroa/internal/interfaces/repository"
)

type IssueService struct {
	issueRepo repository.IssueRepository
	userRepo  repository.UserRepository
}

func NewIssueService(issueRepo repository.IssueRepository, userRepo repository.UserRepository) *IssueService {
	return &IssueService{
		issueRepo: issueRepo,
		userRepo:  userRepo,
	}
}

type CreateIssueRequest struct {
	Title       string
	Description string
	UserID      *uint
}

func (s *IssueService) CreateIssue(req CreateIssueRequest) (*domain.Issue, error) {
	var user *domain.User
	if req.UserID != nil {
		var err error
		user, err = s.userRepo.FindByID(*req.UserID)
		if err != nil {
			return nil, err
		}
	}

	issue := domain.NewIssue(req.Title, req.Description, user)
	if err := s.issueRepo.Create(issue); err != nil {
		return nil, err
	}

	return issue, nil
}

func (s *IssueService) GetIssues(status string) ([]*domain.Issue, error) {
	return s.issueRepo.FindAll(status)
}

func (s *IssueService) GetIssue(id uint) (*domain.Issue, error) {
	return s.issueRepo.FindByID(id)
}

type UpdateIssueRequest struct {
	Title       *string
	Description *string
	Status      *string
	UserID      *uint
}

func (s *IssueService) UpdateIssue(id uint, req UpdateIssueRequest) (*domain.Issue, error) {
	issue, err := s.issueRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	var user *domain.User
	if req.UserID != nil {
		if *req.UserID == 0 {
			user = nil
		} else {
			user, err = s.userRepo.FindByID(*req.UserID)
			if err != nil {
				return nil, err
			}
		}
	}

	title := ""
	if req.Title != nil {
		title = *req.Title
	}

	description := ""
	if req.Description != nil {
		description = *req.Description
	}

	status := ""
	if req.Status != nil {
		status = *req.Status
	}

	if err := issue.Update(title, description, status, user); err != nil {
		return nil, err
	}

	if err := s.issueRepo.Update(issue); err != nil {
		return nil, err
	}

	return issue, nil
} 
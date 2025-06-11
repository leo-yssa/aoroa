package repository

import (
	"aoroa/internal/domain"
	"aoroa/internal/interfaces/repository"
	"aoroa/pkg/common/errors"
	"sync"
)

type MemoryIssueRepository struct {
	issues []*domain.Issue
	mu     sync.RWMutex
}

func NewMemoryIssueRepository() repository.IssueRepository {
	return &MemoryIssueRepository{
		issues: make([]*domain.Issue, 0),
	}
}

func (r *MemoryIssueRepository) Create(issue *domain.Issue) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	issue.ID = uint(len(r.issues) + 1)
	r.issues = append(r.issues, issue)
	return nil
}

func (r *MemoryIssueRepository) FindByID(id uint) (*domain.Issue, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, issue := range r.issues {
		if issue.ID == id {
			return issue, nil
		}
	}
	return nil, errors.ErrIssueNotFound
}

func (r *MemoryIssueRepository) FindAll(status string) ([]*domain.Issue, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if status == "" {
		return r.issues, nil
	}

	filtered := make([]*domain.Issue, 0)
	for _, issue := range r.issues {
		if issue.Status == status {
			filtered = append(filtered, issue)
		}
	}
	return filtered, nil
}

func (r *MemoryIssueRepository) Update(issue *domain.Issue) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, existingIssue := range r.issues {
		if existingIssue.ID == issue.ID {
			r.issues[i] = issue
			return nil
		}
	}
	return errors.ErrIssueNotFound
}

type MemoryUserRepository struct {
	users []*domain.User
	mu    sync.RWMutex
}

func NewMemoryUserRepository() repository.UserRepository {
	users := []*domain.User{
		domain.NewUser(1, "김개발"),
		domain.NewUser(2, "이디자인"),
		domain.NewUser(3, "박기획"),
	}
	return &MemoryUserRepository{
		users: users,
	}
}

func (r *MemoryUserRepository) FindByID(id uint) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.ErrUserNotFound
}

func (r *MemoryUserRepository) FindAll() ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.users, nil
} 
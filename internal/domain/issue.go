package domain

import (
	"errors"
	"time"
)

const (
	StatusPending    = "PENDING"
	StatusInProgress = "IN_PROGRESS"
	StatusCompleted  = "COMPLETED"
	StatusCancelled  = "CANCELLED"
)

var (
	ErrInvalidStatus            = errors.New("invalid status")
	ErrAssigneeRequired        = errors.New("assignee is required for this status")
	ErrIssueCompletedOrCancelled = errors.New("cannot update completed or cancelled issue")
)

var ValidStatuses = map[string]bool{
	StatusPending:    true,
	StatusInProgress: true,
	StatusCompleted:  true,
	StatusCancelled:  true,
}

type Issue struct {
	ID          uint
	Title       string
	Description string
	Status      string
	User        *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewIssue(title, description string, user *User) *Issue {
	now := time.Now()
	status := StatusPending
	if user != nil {
		status = StatusInProgress
	}

	return &Issue{
		Title:       title,
		Description: description,
		Status:      status,
		User:        user,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (i *Issue) Update(title, description string, status string, user *User) error {
	if i.Status == StatusCompleted || i.Status == StatusCancelled {
		return ErrIssueCompletedOrCancelled
	}

	if status != "" {
		if !ValidStatuses[status] {
			return ErrInvalidStatus
		}

		if status != StatusPending && status != StatusCancelled && user == nil {
			return ErrAssigneeRequired
		}
		i.Status = status
	}

	if title != "" {
		i.Title = title
	}
	if description != "" {
		i.Description = description
	}

	if user != nil {
		i.User = user
		if i.Status == StatusPending {
			i.Status = StatusInProgress
		}
	} else if i.User != nil && user == nil {
		i.User = nil
		i.Status = StatusPending
	}

	i.UpdatedAt = time.Now()
	return nil
}

func (i *Issue) IsCompletedOrCancelled() bool {
	return i.Status == StatusCompleted || i.Status == StatusCancelled
} 
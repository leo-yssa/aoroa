package errors

import "errors"

var (
	ErrInvalidStatus            = errors.New("invalid status")
	ErrIssueNotFound           = errors.New("issue not found")
	ErrUserNotFound            = errors.New("user not found")
	ErrInvalidOperation        = errors.New("invalid operation")
	ErrAssigneeRequired        = errors.New("assignee is required for this status")
	ErrIssueCompletedOrCancelled = errors.New("cannot update completed or cancelled issue")
) 
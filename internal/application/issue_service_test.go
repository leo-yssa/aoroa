package application

import (
	"aoroa/internal/domain"
	"aoroa/internal/infrastructure/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupService() *IssueService {
	issueRepo := repository.NewMemoryIssueRepository()
	userRepo := repository.NewMemoryUserRepository()
	return NewIssueService(issueRepo, userRepo)
}

func TestCreateIssue_Success(t *testing.T) {
	svc := setupService()
	title := "테스트 이슈"
	desc := "설명"
	userId := uint(1)
	issue, err := svc.CreateIssue(CreateIssueRequest{
		Title:       title,
		Description: desc,
		UserID:      &userId,
	})
	assert.NoError(t, err)
	assert.Equal(t, title, issue.Title)
	assert.Equal(t, desc, issue.Description)
	assert.Equal(t, domain.StatusInProgress, issue.Status)
	assert.NotNil(t, issue.User)
}

func TestCreateIssue_InvalidUser(t *testing.T) {
	svc := setupService()
	userId := uint(999)
	_, err := svc.CreateIssue(CreateIssueRequest{
		Title:  "테스트",
		UserID: &userId,
	})
	assert.Error(t, err)
}

func TestUpdateIssue_CompletedOrCancelled(t *testing.T) {
	svc := setupService()
	userId := uint(1)
	issue, err := svc.CreateIssue(CreateIssueRequest{
		Title:  "테스트",
		UserID: &userId,
	})
	assert.NoError(t, err)
	assert.NotNil(t, issue)

	// 상태를 COMPLETED로 변경 (담당자가 있는 상태에서)
	completed := domain.StatusCompleted
	updatedIssue, err := svc.UpdateIssue(issue.ID, UpdateIssueRequest{
		Status: &completed,
		UserID: &userId,  // 담당자 유지
	})
	assert.NoError(t, err)
	assert.NotNil(t, updatedIssue)
	assert.Equal(t, domain.StatusCompleted, updatedIssue.Status)

	// 다시 수정 시도
	title := "수정"
	_, err = svc.UpdateIssue(issue.ID, UpdateIssueRequest{Title: &title})
	assert.ErrorIs(t, err, domain.ErrIssueCompletedOrCancelled)
}

func TestGetIssue_NotFound(t *testing.T) {
	svc := setupService()
	_, err := svc.GetIssue(999)
	assert.Error(t, err)
} 
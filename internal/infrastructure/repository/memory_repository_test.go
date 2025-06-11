package repository

import (
	"aoroa/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryUserRepository_FindByID(t *testing.T) {
	repo := NewMemoryUserRepository()
	user, err := repo.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	_, err = repo.FindByID(999)
	assert.Error(t, err)
}

func TestMemoryIssueRepository_CRUD(t *testing.T) {
	repo := NewMemoryIssueRepository()
	issue := &domain.Issue{Title: "테스트", Status: domain.StatusPending}
	err := repo.Create(issue)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), issue.ID)

	found, err := repo.FindByID(issue.ID)
	assert.NoError(t, err)
	assert.Equal(t, issue.Title, found.Title)

	issues, err := repo.FindAll("")
	assert.NoError(t, err)
	assert.Len(t, issues, 1)

	issue.Title = "수정됨"
	err = repo.Update(issue)
	assert.NoError(t, err)
	found, _ = repo.FindByID(issue.ID)
	assert.Equal(t, "수정됨", found.Title)

	// 없는 이슈 업데이트 시 에러
	dummy := &domain.Issue{ID: 999}
	err = repo.Update(dummy)
	assert.Error(t, err)
} 
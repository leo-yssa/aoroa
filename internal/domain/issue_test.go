package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssue_CreateAndUpdate(t *testing.T) {
	user := &User{ID: 1, Name: "김개발"}
	issue := NewIssue("테스트 이슈", "설명", nil)
	assert.Equal(t, StatusPending, issue.Status)
	assert.Nil(t, issue.User)

	// 담당자 할당 시 상태 자동 변경
	err := issue.Update("", "", StatusInProgress, user)
	assert.NoError(t, err)
	assert.Equal(t, StatusInProgress, issue.Status)
	assert.Equal(t, user, issue.User)

	// 상태를 COMPLETED로 변경
	err = issue.Update("", "", StatusCompleted, user)
	assert.NoError(t, err)
	assert.Equal(t, StatusCompleted, issue.Status)

	// COMPLETED 상태에서 수정 시 에러
	err = issue.Update("수정", "", StatusInProgress, user)
	assert.ErrorIs(t, err, ErrIssueCompletedOrCancelled)
}

func TestIssue_InvalidStatus(t *testing.T) {
	user := &User{ID: 1, Name: "김개발"}
	issue := NewIssue("테스트", "", user)

	err := issue.Update("", "", "INVALID", user)
	assert.ErrorIs(t, err, ErrInvalidStatus)
}

func TestIssue_AssigneeRequired(t *testing.T) {
	issue := NewIssue("테스트", "", nil)

	// 담당자 없이 IN_PROGRESS로 변경 시 에러
	err := issue.Update("", "", StatusInProgress, nil)
	assert.ErrorIs(t, err, ErrAssigneeRequired)
}

func TestIssue_RemoveAssignee(t *testing.T) {
	user := &User{ID: 1, Name: "김개발"}
	issue := NewIssue("테스트", "", user)
	assert.Equal(t, StatusInProgress, issue.Status)

	// 담당자 제거 시 상태 PENDING
	err := issue.Update("", "", StatusPending, nil)
	assert.NoError(t, err)
	assert.Nil(t, issue.User)
	assert.Equal(t, StatusPending, issue.Status)
} 
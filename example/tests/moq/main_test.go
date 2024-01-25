package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 方針：
// ExternalAPIClientMockを使って、
// 1.でスタブとして利用し、想定されるユーザーを返却させる。
// 3.でモックとして利用し、保存リクエストを記録・検証する。

// Given-Then-When Pattern(準備・実行・確認)
func TestWithMoq(t *testing.T) {
	// Given

	clientMock := &ExternalAPIClientMock{
		// スタブとして、返却したい値を設定。interfaceの型定義で簡単に定義できる。
		FetchUserFunc: func(id int) (*User, error) {
			return &User{ID: id, Name: "beforeName"}, nil
		},
		UpdateUserFunc: func(user *User) error {
			return nil
		},
	}

	service := NewService(clientMock)

	// When
	err := service.UpdateUserName(1, "afterName")

	// Then: 呼び出し回数が１であること、更新後のユーザー名がafterNameであること
	require.NoError(t, err)
	if assert.Len(t, clientMock.UpdateUserCalls(), 1) {
		// Matcherとして、testifyのエコシステムが使える
		assert.Equal(t, "afterName", clientMock.UpdateUserCalls()[0].User.Name)
	}

}

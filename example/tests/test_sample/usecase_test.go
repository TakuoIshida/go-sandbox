package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubApprover struct {
	isApproved bool
}

func (s *StubApprover) IsApproved() bool {
	return s.isApproved
}

func TestMailSender_CanSend(t *testing.T) {
	tests := []struct {
		name     string
		approver Approver
		email    string
		message  string
		want     bool
	}{
		{
			name:     "未承認の場合は送信できない",
			approver: &StubApprover{isApproved: false},
			want:     false,
		},
		{
			name:     "メールアドレスが空の場合は送信できないこと",
			approver: &StubApprover{isApproved: true},
			want:     false,
		},
		{
			name:     "メッセージが空の場合は送信できないこと",
			approver: &StubApprover{isApproved: true},
			email:    "hoge@example.co.jp",
			want:     false,
		},
		{
			name:     "承認済みでメールアドレスとメッセージがある場合は送信できること",
			approver: &StubApprover{isApproved: true},
			email:    "hoge@example.co.jp",
			message:  "Hello, World!",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := NewMailSender(tt.approver, tt.email)
			sender.SetMessage(tt.message)
			got := sender.CanSend()

			assert.Equal(t, tt.want, got)
		})
	}

}

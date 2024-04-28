package service

import (
	"errors"
	"go-sandbox/example/failure/apperror"

	"github.com/morikuni/failure"
)

type User struct{}

func (h User) Search(id string) (string, error) {
	// 処理...
	err := errors.New("not found error")
	if err != nil {
		// err を wrap してエラー情報を追加する
		return "", failure.Translate(err, apperror.ClientError, failure.Messagef("invalid id=%v", id))
	}
	return "user found", nil
}

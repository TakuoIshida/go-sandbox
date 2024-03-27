package usecases

import (
	"reflect"
	"testing"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		userRepo UserRepo
	}
	tests := []struct {
		name string
		args args
		want UserUsecase
	}{
		// テストケース2: nil のユーザーリポジトリを渡す
		{
			name: "nilのユーザーリポジトリを渡す",
			args: args{
				userRepo: nil,
			},
			want: UserUsecase{},
		},
	}
	t.Fatalf("Fatal")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_GetMe(t *testing.T) {
	type fields struct {
		userRepo UserRepo
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserUsecase{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.GetMe(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.GetMe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.GetMe() = %v, want %v", got, tt.want)
			}
		})
	}
}

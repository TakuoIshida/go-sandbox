package main

// moq ライブラリを使って、IFからそのmock, stubを生成する
// https://zenn.dev/abekoh/articles/21acde07e1f555#%E3%83%A2%E3%83%83%E3%82%AF%E3%81%A8%E3%82%B9%E3%82%BF%E3%83%96

// moq -out main_moq.go . ExternalAPIClient
type User struct {
	ID   int
	Name string
}

type ExternalAPIClient interface {
	FetchUser(id int) (*User, error)
	UpdateUser(user *User) error
}

type Service struct {
	client ExternalAPIClient
}

func NewService(client ExternalAPIClient) *Service {
	return &Service{client: client}
}

// UpdateUserName は与えられた idのユーザー名を与えられた nameに更新する
// 1. ExternalAPIClient.FetchUser()でidのユーザーを取得
// 2. ユーザー名をnameに置き換え
// 3. ExternalAPIClient.UpdateUser()でユーザーを保存
func (s Service) UpdateUserName(id int, name string) error {
	user, err := s.client.FetchUser(id)
	if err != nil {
		return err
	}

	user.Name = name
	return s.client.UpdateUser(user)
}

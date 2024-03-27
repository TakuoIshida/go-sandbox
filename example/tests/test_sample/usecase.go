package usecases

type (
	UserUsecase struct {
		userRepo UserRepo
	}

	UserRepo interface {
		Find(id int) User
	}

	User struct {
		id    string
		name  string
		email string
	}
)

func NewUserUsecase(userRepo UserRepo) UserUsecase {
	return UserUsecase{userRepo}
}

func (u UserUsecase) GetMe(id int) (User, error) {
	return u.userRepo.Find(id), nil
}

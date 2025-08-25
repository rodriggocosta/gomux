package users

import (
	"apigo/entity"
	"apigo/repository/users"
)

type UserGetUsecase struct {
	repository users.UsersGetRepository
}

func NewUserGet(repo users.UsersGetRepository) UserGetUsecase {
	return UserGetUsecase{
		repository: repo,
	}
}

func (ur *UserGetUsecase) GetUser() ([]entity.Users, error) {
	return ur.repository.GetUser()
}

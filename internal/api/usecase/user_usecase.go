package usecase

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/repository"

	bycrypt "golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.UserRequest) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(u model.UserRequest) (string, error) {

	hash,err := bycrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return "", err
	}

	newUser := model.UserRequest{
		Name: u.Name,
		Email: u.Email,
		Password: string(hash),
	}

	res,err :=uu.ur.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return res, nil
}
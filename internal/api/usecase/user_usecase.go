package usecase

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/repository"
)

type IUserUsecase interface {
	SignUp(req model.UserRequest) (string, error)
	Login(model.UserRequest) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(req model.UserRequest) (string, error) {

	newUser := model.UserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := uu.ur.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (uu *userUsecase) Login(req model.UserRequest) (string, error) {

	newUser := model.UserRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := uu.ur.LoginUser(newUser)
	if err != nil {
		return "", err
	}

	return res, nil
}



package user_service

import (
	"net/http"
	"strconv"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/pkg/helpers"
	"toko-belanja-app/repository/user_repository"
)

type UserService interface {
	RegisterUser(userPayLoad *dto.CreateNewUsersRequest) (*dto.UserResponse, errs.Error)
	LoginUser(userPayLoad *dto.UsersLoginRequest) (*dto.UserResponse, errs.Error)
	TopUpBalance(userId int, userPayLoad *dto.UsersTopUpRequest) (*dto.UserResponse, errs.Error)
}

type userServiceImpl struct {
	ur user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userServiceImpl{ur: userRepo}
}

func (us *userServiceImpl) RegisterUser(userPayLoad *dto.CreateNewUsersRequest) (*dto.UserResponse, errs.Error) {
	err := helpers.ValidateStruct(userPayLoad)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		FullName: userPayLoad.FullName,
		Email:    userPayLoad.Email,
		Password: userPayLoad.Password,
	}

	user.HashPassword()

	response, err := us.ur.CreateNewUser(user)

	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Code:    http.StatusCreated,
		Message: "Your account has been successfully created",
		Data:    response,
	}, nil
}



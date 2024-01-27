package service

import (
	"errors"
	"github.com/adibhauzan/master_data/helper"
	"github.com/adibhauzan/master_data/model/domain"
	"github.com/adibhauzan/master_data/model/web/request"
	"github.com/adibhauzan/master_data/model/web/response"
	"github.com/adibhauzan/master_data/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx *gin.Context, request request.CreateUserRequest) (response.CreateUserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return response.CreateUserResponse{}, errors.New("Bad Request Validasinya yang bener dong tolol")
	}

	var user domain.User
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		user = domain.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
		}

		user, err = service.UserRepository.Save(ctx, user)
		if err != nil {
			panic(err)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user), nil
}

package repository

import (
	"github.com/adibhauzan/master_data/model/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func hashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	hashedPasswordString := string(hashedPassword)
	return hashedPasswordString, nil
}

func (repository *UserRepositoryImpl) Save(ctx *gin.Context, user domain.User) (domain.User, error) {

	userPassword, err := hashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	dbUser := domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: userPassword,
	}

	err = repository.DB.Create(&dbUser).Error
	if err != nil {
		panic(err)
		return domain.User{}, err
	}

	return dbUser, nil
}

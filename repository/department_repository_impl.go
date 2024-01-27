package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/adibhauzan/master_data/model/domain"
	"gorm.io/gorm"
)

type DepartmentRepositoryImpl struct {
	DB *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &DepartmentRepositoryImpl{
		DB: db,
	}
}

func (repository *DepartmentRepositoryImpl) Save(ctx *gin.Context, department domain.Department) domain.Department {
	err := repository.DB.Create(&department).Error
	if err != nil {
		panic(err)
	}

	return department
}

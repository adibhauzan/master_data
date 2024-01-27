package service

import (
	"github.com/adibhauzan/master_data/helper"
	"github.com/adibhauzan/master_data/model/domain"
	"github.com/adibhauzan/master_data/model/web/request"
	"github.com/adibhauzan/master_data/model/web/response"
	"github.com/adibhauzan/master_data/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type DepartmentServiceImpl struct {
	DepartmentRepository repository.DepartmentRepository
	DB                   *gorm.DB
	Validate             *validator.Validate
}

func NewDepartmentService(departmentRepo repository.DepartmentRepository, db *gorm.DB, validate *validator.Validate) DepartmentService {
	return &DepartmentServiceImpl{
		DepartmentRepository: departmentRepo,
		DB:                   db,
		Validate:             validate,
	}
}
func (service *DepartmentServiceImpl) Create(ctx *gin.Context, request request.DepartmentRequest) response.CreateDepartmentResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}

	var department domain.Department
	err = service.DB.Transaction(func(tx *gorm.DB) error {
		department = domain.Department{
			Name: request.Name,
		}

		department = service.DepartmentRepository.Save(ctx, department)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return helper.ToDepartmentResponse(department)
}

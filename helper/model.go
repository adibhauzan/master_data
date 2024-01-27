package helper

import (
	"github.com/adibhauzan/master_data/model/domain"
	"github.com/adibhauzan/master_data/model/web/response"
)

func ToDepartmentResponse(department domain.Department) response.CreateDepartmentResponse {
	return response.CreateDepartmentResponse{
		ID:        department.ID,
		Name:      department.Name,
		CreatedAt: department.CreatedAt,
		UpdatedAt: department.UpdatedAt,
	}
}

func ToUserResponse(user domain.User) response.CreateUserResponse {
	return response.CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

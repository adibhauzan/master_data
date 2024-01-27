package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/adibhauzan/master_data/model/domain"
)

type DepartmentRepository interface {
	Save(ctx *gin.Context, department domain.Department) domain.Department
}

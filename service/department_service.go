package service

import (
	"github.com/adibhauzan/master_data/model/web/request"
	"github.com/adibhauzan/master_data/model/web/response"
	"github.com/gin-gonic/gin"
)

type DepartmentService interface {
	Create(ctx *gin.Context, request request.DepartmentRequest) response.CreateDepartmentResponse
}

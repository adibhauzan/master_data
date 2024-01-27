package controller

import (
	"github.com/adibhauzan/master_data/model/web/request"
	"github.com/adibhauzan/master_data/model/web/response"
	"github.com/adibhauzan/master_data/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DepartmentControllerImpl struct {
	DepartmentService service.DepartmentService
}

func NewDepartmentController(departmentService service.DepartmentService) DepartmentController {
	return &DepartmentControllerImpl{
		DepartmentService: departmentService,
	}
}

func (controller *DepartmentControllerImpl) Create(ctx *gin.Context) {
	var departmentRequest request.DepartmentRequest

	if err := ctx.ShouldBindJSON(&departmentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	departmentResponse := controller.DepartmentService.Create(ctx, departmentRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success Create Department",
		Data:   departmentResponse,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

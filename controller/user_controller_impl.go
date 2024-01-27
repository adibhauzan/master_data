package controller

import (
	"github.com/adibhauzan/master_data/model/web/request"
	"github.com/adibhauzan/master_data/model/web/response"
	"github.com/adibhauzan/master_data/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(ctx *gin.Context) {
	var userRequest request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userResponse, err := controller.UserService.Create(ctx, userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error ngab": err.Error()})
	}

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success Create User",
		Data:   userResponse,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

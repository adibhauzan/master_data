package controller

import "github.com/gin-gonic/gin"

type DepartmentController interface {
	Create(ctx *gin.Context)
}

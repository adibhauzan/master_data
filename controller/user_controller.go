package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Create(ctx *gin.Context)
}

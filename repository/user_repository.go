package repository

import (
	"github.com/adibhauzan/master_data/model/domain"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Save(ctx *gin.Context, department domain.User) (domain.User, error)
}

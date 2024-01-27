package main

import (
	"github.com/adibhauzan/master_data/app"
	"github.com/adibhauzan/master_data/controller"
	"github.com/adibhauzan/master_data/model/domain"
	"github.com/adibhauzan/master_data/repository"
	"github.com/adibhauzan/master_data/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	db.AutoMigrate(&domain.Department{}, &domain.User{})
	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	departmentRepository := repository.NewDepartmentRepository(db)
	departmentService := service.NewDepartmentService(departmentRepository, db, validate)
	departmentController := controller.NewDepartmentController(departmentService)

	router := app.NewRouter(departmentController, userController)

	err := router.Run(":3000")
	if err != nil {
		return
	}
}

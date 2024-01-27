package app

import (
	"github.com/adibhauzan/master_data/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(departmentController controller.DepartmentController, userController controller.UserController) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	departmentRoutes(api, departmentController)
	userRoutes(api, userController)
	return r
}

func departmentRoutes(api *gin.RouterGroup, departmentController controller.DepartmentController) {
	departmentApi := api.Group("/department")
	{
		departmentApi.POST("/", departmentController.Create)
	}
}

func userRoutes(api *gin.RouterGroup, userController controller.UserController) {
	userApi := api.Group("/user")
	{
		userApi.POST("/", userController.Create)
	}
}

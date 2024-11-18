package routes

import (
	"employee-management-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(employeeController *controllers.EmployeeController) *gin.Engine {
	router := gin.Default()

	router.GET("/employees", employeeController.GetAllEmployees)
	router.GET("/employees/:id", employeeController.GetEmployeeByID)
	router.POST("/employees", employeeController.AddEmployee)
	router.PUT("/employees/:id", employeeController.UpdateEmployee)
	router.DELETE("/employees/:id", employeeController.DeleteEmployee)

	return router
}

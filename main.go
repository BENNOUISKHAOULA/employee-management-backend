package main

import (
	"employee-management-backend/config"
	"employee-management-backend/controllers"
	"employee-management-backend/repositories"
	"employee-management-backend/routes"
	"employee-management-backend/services"
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Setup dependencies
	employeeRepo := repositories.NewEmployeeRepository(config.GetCollection("employees"))
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	// Setup routes
	router := routes.SetupRouter(employeeController)

	// Start server
	router.Run(":8080")
}

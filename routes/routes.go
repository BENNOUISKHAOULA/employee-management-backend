package routes

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/employees", GetEmployees)
    router.POST("/employees", AddEmployee)
    router.PUT("/employees/:id", UpdateEmployee)
    router.DELETE("/employees/:id", DeleteEmployee)
    return router
}

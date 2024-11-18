package controllers

import (
	"employee-management-backend/models"
	"employee-management-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (ctl *EmployeeController) GetAllEmployees(c *gin.Context) {
	employees, err := ctl.service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (ctl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	employee, err := ctl.service.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (ctl *EmployeeController) AddEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := ctl.service.AddEmployee(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (ctl *EmployeeController) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.service.UpdateEmployee(id, &employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

func (ctl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	err := ctl.service.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

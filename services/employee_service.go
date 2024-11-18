package services

import (
	"employee-management-backend/models"
	"employee-management-backend/repositories"
)

type EmployeeService interface {
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	AddEmployee(employee *models.Employee) (string, error)
	UpdateEmployee(id string, employee *models.Employee) error
	DeleteEmployee(id string) error
}

type employeeService struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.FindAll()
}

func (s *employeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *employeeService) AddEmployee(employee *models.Employee) (string, error) {
	return s.repo.Create(employee)
}

func (s *employeeService) UpdateEmployee(id string, employee *models.Employee) error {
	return s.repo.Update(id, employee)
}

func (s *employeeService) DeleteEmployee(id string) error {
	return s.repo.Delete(id)
}

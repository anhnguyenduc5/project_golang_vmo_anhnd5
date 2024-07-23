package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/storage"
)

// DepartmentUsecase handles the business logic for departments.
type DepartmentUsecase struct {
	repo *storage.DepartmentRepository
}

// NewDepartmentUsecase creates a new instance of DepartmentUsecase.
func NewDepartmentUsecase(repo *storage.DepartmentRepository) *DepartmentUsecase {
	return &DepartmentUsecase{repo: repo}
}

// CreateDepartment creates a new department using the provided DTO.
func (u *DepartmentUsecase) CreateDepartment(input dto.DepartmentCreateDTO) (*domain.Department, error) {
	department := &domain.Department{
		Name:    input.Name,
		Address: input.Address,
		Status:  input.Status,
	}
	err := u.repo.Create(department)
	return department, err
}

func (u *DepartmentUsecase) GetAllDepartments() ([]domain.Department, error) {
	return u.repo.GetAll()
}

// GetDepartmentByID retrieves a department by its ID.
func (u *DepartmentUsecase) GetDepartmentByID(id uint) (*domain.Department, error) {
	return u.repo.GetByID(id)
}

// UpdateDepartment updates a department using the provided DTO.
func (u *DepartmentUsecase) UpdateDepartment(id uint, input dto.DepartmentUpdateDTO) (*domain.Department, error) {
	department, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	department.Name = input.Name
	department.Address = input.Address
	department.Status = input.Status
	err = u.repo.Update(department)
	return department, err
}

// DeleteDepartment deletes a department by its ID.
func (u *DepartmentUsecase) DeleteDepartment(id uint) error {
	return u.repo.Delete(id)
}

package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/usecase"
	"github.com/gin-gonic/gin"
)

// DepartmentHandler handles the HTTP requests for departments.
type DepartmentHandler struct {
	usecase *usecase.DepartmentUsecase
}

// NewDepartmentHandler creates a new instance of DepartmentHandler.
func NewDepartmentHandler(usecase *usecase.DepartmentUsecase) *DepartmentHandler {
	return &DepartmentHandler{usecase: usecase}
}

// CreateDepartment handles the HTTP POST request to create a new department.
// CreateDepartment godoc
// @Summary Create a new department
// @Description Create a new department
// @Accept json
// @Produce json
// @Tags department
// @Param department body dto.DepartmentCreateDTO true "Department data"
// @Success 201 {object} domain.Department
// @Router /api/v1/departments [post]
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var input dto.DepartmentCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.usecase.CreateDepartment(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)
}

// GetAllDepartments handles the HTTP GET request to retrieve all departments.
// GetAllDepartments godoc
// @Summary Get all departments
// @Description Get all departments
// @Produce json
// @Tags department
// @Success 200 {array} domain.Department
// @Router /api/v1/departments [get]
func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	departments, err := h.usecase.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, departments)
}

// GetDepartmentByID handles the HTTP GET request to retrieve a department by its ID.
// GetDepartmentByID godoc
// @Summary Get department by ID
// @Description Get department by ID
// @Produce json
// @Tags department
// @Param id path int true "Department ID"
// @Success 200 {object} domain.Department
// @Router /api/v1/departments/{id} [get]
func (h *DepartmentHandler) GetDepartmentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	department, err := h.usecase.GetDepartmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	c.JSON(http.StatusOK, department)
}

// UpdateDepartment handles the HTTP PUT request to update a department.
// UpdateDepartment godoc
// @Summary Update department
// @Description Update department
// @Accept json
// @Produce json
// @Tags department
// @Param id path int true "Department ID"
// @Param department body dto.DepartmentUpdateDTO true "Department data"
// @Success 200 {object} domain.Department
// @Router /api/v1/departments/{id} [put]
func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	var input dto.DepartmentUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.usecase.UpdateDepartment(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

// DeleteDepartment handles the HTTP DELETE request to delete a department.
// DeleteDepartment godoc
// @Summary Delete department
// @Description Delete department
// @Produce json
// @Tags department
// @Param id path int true "Department ID"
// @Success 204
// @Router /api/v1/departments/{id} [delete]
func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	err = h.usecase.DeleteDepartment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

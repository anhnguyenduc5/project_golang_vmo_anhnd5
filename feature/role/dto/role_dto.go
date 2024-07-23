package dto

// RoleCreateDTO represents the data transfer object for creating a role.
type RoleCreateDTO struct {
	Name   string `json:"name" binding:"required"`
	Status uint   `json:"status" binding:"required"`
}

// RoleUpdateDTO represents the data transfer object for updating a role.
type RoleUpdateDTO struct {
	Name   string `json:"name" binding:"required"`
	Status uint   `json:"status" binding:"required"`
}

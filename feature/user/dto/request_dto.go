package dto

type RequestCreatingDTO struct {
	UserID            int     `json:"user_id" binding:"required"`
	DepartmentID      *int    `json:"department_id" binding:"required"`
	Gender            *string `json:"gender"`
	DOB               string  `json:"dob"`
	Mobile            *string `json:"mobile"`
	CountryID         *int    `json:"country_id"`
	ResidentCountryID *int    `json:"resident_country_id"`
}

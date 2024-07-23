package storage

import (
	"errors"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

	"gorm.io/gorm"
)

type ApplicantRequestRepositoryInterface interface {
	CreateApplicantRequest(reqRequest *domain.Request, reqUser *domain.User) error
}

type ApplicantRequestRepository struct {
	db *gorm.DB
}

func NewApplicantRequestRepository(db *gorm.DB) *ApplicantRequestRepository {
	return &ApplicantRequestRepository{db: db}
}

func (r *ApplicantRequestRepository) CreateApplicantRequest(reqRequest *domain.Request, reqUser *domain.User) error {
	// find request
	var existingRequests []domain.Request
	query := r.db.Where("user_id = ?", reqUser.ID).Find(&existingRequests)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected > 0 {
		return errors.New("this user already has a request")
	}
	//find user
	if err := r.db.First(&domain.User{}, reqUser.ID).Error; err != nil {
		return errors.New("user not found")
	}
	// update user
	result := r.db.Model(&domain.User{}).Where("id = ?", reqUser.ID).Updates(map[string]interface{}{
		"department_id":       reqUser.DepartmentID,
		"gender":              reqUser.Gender,
		"dob":                 reqUser.Dob,
		"mobile":              reqUser.Mobile,
		"country_id":          reqUser.CountryID,
		"resident_country_id": reqUser.ResidentCountryID,
	})
	if result.Error != nil {
		return result.Error
	}
	return r.db.Create(reqRequest).Error
}

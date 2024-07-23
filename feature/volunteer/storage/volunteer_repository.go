package storage

import (
	"gorm.io/gorm"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
)

// VolunteerRepositoryInterface defines the methods that a VolunteerRepository should implement
type VolunteerRepositoryInterface interface {
	CreateVolunteer(volunteer *domain.VolunteerDetails) error
	UpdateVolunteer(volunteer *domain.VolunteerDetails) error
	DeleteVolunteer(id int) error
	FindVolunteerByID(id int) (*domain.VolunteerDetails, error)
	GetAllVolunteers() ([]*domain.VolunteerDetails, error)
}

type VolunteerRepository struct {
	db *gorm.DB
}

func NewVolunteerRepository(db *gorm.DB) *VolunteerRepository {
	return &VolunteerRepository{db: db}
}

func (r *VolunteerRepository) CreateVolunteer(volunteer *domain.VolunteerDetails) error {
	return r.db.Create(volunteer).Error
}

func (r *VolunteerRepository) UpdateVolunteer(volunteer *domain.VolunteerDetails) error {
	return r.db.Save(volunteer).Error
}

func (r *VolunteerRepository) DeleteVolunteer(id int) error {
	return r.db.Delete(&domain.VolunteerDetails{}, id).Error
}

func (r *VolunteerRepository) FindVolunteerByID(id int) (*domain.VolunteerDetails, error) {
	var volunteer *domain.VolunteerDetails
	if err := r.db.First(&volunteer, id).Error; err != nil {
		return nil, err
	}
	return volunteer, nil
}

func (r *VolunteerRepository) GetAllVolunteers() ([]*domain.VolunteerDetails, error) {
	var volunteers []*domain.VolunteerDetails
	if err := r.db.Find(&volunteers).Error; err != nil {
		return nil, err
	}
	return volunteers, nil
}

package usecase

import (
	"errors"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type VolunteerRequestUsecaseInterface interface {
	CreateVolunteerRequest(request dto.RequestCreatingDTO) error
}

type VolunteerRequestUsecase struct {
	VolRequestRepo storage.VolunteerRequestRepositoryInterface
}

func NewVolunteerRequestUsecase(volRequestRepo storage.VolunteerRequestRepositoryInterface) *VolunteerRequestUsecase {
	return &VolunteerRequestUsecase{VolRequestRepo: volRequestRepo}
}

func (u *VolunteerRequestUsecase) CreateVolunteerRequest(request dto.RequestCreatingDTO) error {
	err := ValidateInput(request)
	if err != nil {
		return err
	}
	reqRequest := &domain.Request{
		UserID:     request.UserID,
		Type:       "verification",
		Status:     0,
		VerifierID: nil,
	}
	parsedTime, err := StringToTimePtr(request.DOB)
	if err != nil {
		return errors.New("invalid date of birth")
	}
	roleID := 2
	reqUser := &domain.User{
		ID:                request.UserID,
		DepartmentID:      request.DepartmentID,
		Gender:            request.Gender,
		Dob:               parsedTime,
		Mobile:            request.Mobile,
		CountryID:         request.CountryID,
		ResidentCountryID: request.ResidentCountryID,
		RoleID:            &roleID,
	}
	return u.VolRequestRepo.CreateVolunteerRequest(reqRequest, reqUser)
}

func ValidateInput(request dto.RequestCreatingDTO) error {
	genderMap := map[string]bool{
		"Male":   true,
		"Female": true,
		"male":   true,
		"female": true,
		"Other":  true,
		"other":  true,
	}
	if request.Gender == nil || !genderMap[*request.Gender] {
		return errors.New("invalid gender")
	}
	if request.Mobile == nil || len(*request.Mobile) != 10 || (*request.Mobile)[0] != '0' {
		return errors.New("invalid mobile number")
	}
	return nil
}

package usecase

import (
	"errors"

	"github.com/silence99999/doctor_service/internal/model"
	"github.com/silence99999/doctor_service/internal/repository"

	"github.com/google/uuid"
)

type DoctorUsecase struct {
	repo repository.DoctorRepository
}

func NewDoctorUsecase(r repository.DoctorRepository) *DoctorUsecase {
	return &DoctorUsecase{repo: r}
}

func (u *DoctorUsecase) CreateDoctor(fullName, specialization, email string) (model.Doctor, error) {
	if fullName == "" {
		return model.Doctor{}, errors.New("full name is required")
	}
	if email == "" {
		return model.Doctor{}, errors.New("email is required")
	}

	doc := model.Doctor{
		ID:             uuid.NewString(),
		FullName:       fullName,
		Specialization: specialization,
		Email:          email,
	}

	err := u.repo.Create(doc)
	return doc, err
}

func (u *DoctorUsecase) GetDoctor(id string) (model.Doctor, error) {
	return u.repo.GetByID(id)
}

func (u *DoctorUsecase) GetAllDoctors() ([]model.Doctor, error) {
	return u.repo.GetAll()
}

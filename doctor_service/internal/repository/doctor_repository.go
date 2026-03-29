package repository

import "github.com/silence99999/doctor_service/internal/model"

type DoctorRepository interface {
	Create(doctor model.Doctor) error
	GetByID(id string) (model.Doctor, error)
	GetAll() ([]model.Doctor, error)
	GetByEmail(email string) (model.Doctor, error)
}

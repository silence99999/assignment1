package usecase

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/silence99999/appointment_service/internal/client"
	"github.com/silence99999/appointment_service/internal/model"
	"github.com/silence99999/appointment_service/internal/repository"
)

type AppointmentUsecase struct {
	repo   repository.AppointmentRepository
	doctor client.DoctorClient
}

func NewAppointmentUsecase(repo repository.AppointmentRepository, doctor client.DoctorClient) *AppointmentUsecase {
	return &AppointmentUsecase{
		repo:   repo,
		doctor: doctor,
	}
}

func (u *AppointmentUsecase) CreateAppointment(title, description, doctorID string) (model.Appointment, error) {
	if title == "" {
		return model.Appointment{}, errors.New("title is required")
	}
	if doctorID == "" {
		return model.Appointment{}, errors.New("doctor id is required")
	}

	exists, err := u.doctor.Exists(doctorID)
	if err != nil {
		return model.Appointment{}, errors.New("doctor service is unavailable")
	}
	if !exists {
		return model.Appointment{}, errors.New("there is no doctor with this id")
	}

	appointment := model.Appointment{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		DoctorID:    doctorID,
		Status:      model.StatusNew,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = u.repo.Create(appointment)
	return appointment, err
}

func (u *AppointmentUsecase) GetAppointment(id string) (model.Appointment, error) {
	return u.repo.GetByID(id)
}

func (u *AppointmentUsecase) GetAllAppointments() ([]model.Appointment, error) {
	return u.repo.GetAll()
}

func (u *AppointmentUsecase) UpdateAppointmentStatus(newStatus model.Status, id string) error {
	if !newStatus.IsValid() {
		return errors.New("invalid status")
	}

	appointment, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	if appointment.Status == model.StatusDone && newStatus == model.StatusNew {
		return errors.New("cannot change status from done to new")
	}

	updatedAt := time.Now()

	err = u.repo.UpdateStatusByID(newStatus, updatedAt, id)
	return err
}

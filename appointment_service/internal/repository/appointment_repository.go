package repository

import (
	"time"

	"github.com/silence99999/appointment_service/internal/model"
)

type AppointmentRepository interface {
	Create(appointment model.Appointment) error
	GetByID(id string) (model.Appointment, error)
	GetAll() ([]model.Appointment, error)
	UpdateStatusByID(status model.Status, updated_at time.Time, id string) error
}

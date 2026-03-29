package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/silence99999/appointment_service/internal/model"
)

type PostgresAppointmentrRepo struct {
	db *pgxpool.Pool
}

func NewPostgresAppointmentRepo(db *pgxpool.Pool) *PostgresAppointmentrRepo {
	return &PostgresAppointmentrRepo{
		db: db,
	}
}

func (p *PostgresAppointmentrRepo) Create(appointment model.Appointment) error {
	query := `INSERT INTO appointments (id,title,description,doctor_id,status,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	args := []interface{}{appointment.ID, appointment.Title, appointment.Description, appointment.DoctorID, appointment.Status, appointment.CreatedAt, appointment.UpdatedAt}

	_, err := p.db.Exec(ctx, query, args...)

	return err
}

func (p *PostgresAppointmentrRepo) GetByID(id string) (model.Appointment, error) {
	query := `SELECT id,title,description, doctor_id,status,created_at,updated_at FROM appointments WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var appointment model.Appointment

	err := p.db.QueryRow(ctx, query, id).Scan(&appointment.ID, &appointment.Title, &appointment.Description, &appointment.DoctorID, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)

	if err != nil {
		return model.Appointment{}, errors.New("appointment not found")
	}

	return appointment, nil
}

func (p *PostgresAppointmentrRepo) GetAll() ([]model.Appointment, error) {
	query := `SELECT id,title,description,doctor_id,status,created_at,updated_at FROM appointments`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []model.Appointment

	for rows.Next() {
		var appointment model.Appointment
		err := rows.Scan(&appointment.ID, &appointment.Title, &appointment.Description, &appointment.DoctorID, &appointment.Status, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (p *PostgresAppointmentrRepo) UpdateStatusByID(status model.Status, updatedAt time.Time, id string) error {
	query := `UPDATE appointments SET status = $1,updated_at = $2 WHERE id = $3`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	args := []interface{}{status, updatedAt, id}

	_, err := p.db.Exec(ctx, query, args...)

	return err
}

package repository

import (
	"context"
	"errors"
	"time"

	"github.com/silence99999/doctor_service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDoctorRepo struct {
	db *pgxpool.Pool
}

func NewPostgresDoctorRepo(db *pgxpool.Pool) *PostgresDoctorRepo {
	return &PostgresDoctorRepo{db: db}
}

func (r *PostgresDoctorRepo) Create(doc model.Doctor) error {
	query := `
		INSERT INTO doctors (id, full_name, specialization, email)
		VALUES ($1, $2, $3, $4)
	`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	args := []interface{}{doc.ID, doc.FullName, doc.Specialization, doc.Email}

	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *PostgresDoctorRepo) GetByID(id string) (model.Doctor, error) {
	query := `
		SELECT id, full_name, specialization, email
		FROM doctors
		WHERE id = $1
	`

	var doc model.Doctor

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := r.db.QueryRow(ctx, query, id).
		Scan(&doc.ID, &doc.FullName, &doc.Specialization, &doc.Email)

	if err != nil {
		return model.Doctor{}, errors.New("doctor not found")
	}

	return doc, nil
}

func (r *PostgresDoctorRepo) GetAll() ([]model.Doctor, error) {
	query := `
		SELECT id, full_name, specialization, email
		FROM doctors
	`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []model.Doctor

	for rows.Next() {
		var doc model.Doctor
		err := rows.Scan(&doc.ID, &doc.FullName, &doc.Specialization, &doc.Email)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doc)
	}

	return doctors, nil
}

func (r *PostgresDoctorRepo) GetByEmail(email string) (model.Doctor, error) {
	query := `
		SELECT id, full_name, specialization, email
		FROM doctors
		WHERE email = $1
	`

	var doc model.Doctor

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := r.db.QueryRow(ctx, query, email).
		Scan(&doc.ID, &doc.FullName, &doc.Specialization, &doc.Email)

	if err != nil {
		return model.Doctor{}, errors.New("not found")
	}

	return doc, nil
}

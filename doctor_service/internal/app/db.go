package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() *pgxpool.Pool {
	connStr := os.Getenv("DOCTOR_DB_DSN")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal("cannot create pool:", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	log.Println("successfully connected to db")

	return db
}

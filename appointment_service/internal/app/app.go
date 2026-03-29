package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/silence99999/appointment_service/internal/client"
	"github.com/silence99999/appointment_service/internal/repository"
	"github.com/silence99999/appointment_service/internal/transport/http"
	"github.com/silence99999/appointment_service/internal/usecase"
)

func Run() {
	godotenv.Load()
	db := NewDB()

	doctorBaseURL := os.Getenv("DOCTOR_SERVICE_URL")

	doctorClient := client.NewHTTPDoctorClient(doctorBaseURL)

	repo := repository.NewPostgresAppointmentRepo(db)
	usecase := usecase.NewAppointmentUsecase(repo, doctorClient)
	handler := http.NewAppointmentHandler(usecase)

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8081")
}

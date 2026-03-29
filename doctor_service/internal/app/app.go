package app

import (
	"github.com/joho/godotenv"
	"github.com/silence99999/doctor_service/internal/repository"
	"github.com/silence99999/doctor_service/internal/transport/http"
	"github.com/silence99999/doctor_service/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Run() {
	godotenv.Load()
	db := NewDB()

	repo := repository.NewPostgresDoctorRepo(db)
	usecase := usecase.NewDoctorUsecase(repo)
	handler := http.NewDoctorHandler(usecase)

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8080")
}

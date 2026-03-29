package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silence99999/appointment_service/internal/model"
	"github.com/silence99999/appointment_service/internal/usecase"
)

type AppointmentHandler struct {
	usecase *usecase.AppointmentUsecase
}

func NewAppointmentHandler(usecase *usecase.AppointmentUsecase) *AppointmentHandler {
	return &AppointmentHandler{
		usecase: usecase,
	}
}

func (h *AppointmentHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/appointments", h.CreateAppointment)
	r.GET("/appointments/:id", h.GetAppointment)
	r.GET("/appointments", h.GetAllAppointments)
	r.PATCH("/appointments/:id/status", h.UpdateAppointmentStatus)
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DoctorID    string `json:"doctor_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment, err := h.usecase.CreateAppointment(req.Title, req.Description, req.DoctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
	id := c.Param("id")

	appointment, err := h.usecase.GetAppointment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	appointments, err := h.usecase.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func (h *AppointmentHandler) UpdateAppointmentStatus(c *gin.Context) {
	var req struct {
		Status model.Status `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	err := h.usecase.UpdateAppointmentStatus(req.Status, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

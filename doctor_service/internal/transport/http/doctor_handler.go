package http

import (
	"net/http"

	"github.com/silence99999/doctor_service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	usecase *usecase.DoctorUsecase
}

func NewDoctorHandler(u *usecase.DoctorUsecase) *DoctorHandler {
	return &DoctorHandler{usecase: u}
}

func (h *DoctorHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/doctors", h.CreateDoctor)
	r.GET("/doctors/:id", h.GetDoctor)
	r.GET("/doctors", h.GetAllDoctors)
}

func (h *DoctorHandler) CreateDoctor(c *gin.Context) {
	var req struct {
		FullName       string `json:"full_name"`
		Specialization string `json:"specialization"`
		Email          string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doc, err := h.usecase.CreateDoctor(req.FullName, req.Specialization, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, doc)
}

func (h *DoctorHandler) GetDoctor(c *gin.Context) {
	id := c.Param("id")

	doc, err := h.usecase.GetDoctor(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doc)
}

func (h *DoctorHandler) GetAllDoctors(c *gin.Context) {
	docs, err := h.usecase.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, docs)
}

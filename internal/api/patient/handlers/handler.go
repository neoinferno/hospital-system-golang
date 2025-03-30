package handlers

import (
	hospital "hospital-system/internal/api/patient"
)

type Handler struct {
	service hospital.Service
}

func NewHandlers(service hospital.Service) *Handler {
	return &Handler{service: service}
}

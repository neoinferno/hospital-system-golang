package handlers

import (
	hospital "hospital-system/internal/api/staff"
)

type Handler struct {
	service hospital.Service
}

func NewHandlers(service hospital.Service) *Handler {
	return &Handler{service: service}
}

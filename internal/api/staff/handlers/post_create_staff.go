package handlers

import (
	"hospital-system/internal/api/staff/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostCreateStaff(c *gin.Context) {
	var request models.PostCreateStaff
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.PostCreateStaff(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"status":  "success",
		"message": "Created staff success",
	})
}

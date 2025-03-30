package handlers

import (
	"hospital-system/internal/api/staff/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostLogin(c *gin.Context) {
	var request models.LoginRequest

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.PostLogin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    token,
		"status":  "success",
		"message": "Login success",
	})
}

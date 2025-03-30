package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPatients(c *gin.Context) {

	hospitalId, exists := c.Get("hospitalID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "can't get hospitalID"})
		return
	}
	filters := buildFilter(c, hospitalId.(string))

	patients, err := h.service.GetAllPatients(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    patients,
		"status":  "success",
		"message": "Get all patients success",
	})
}

func buildFilter(c *gin.Context, hospitalId string) map[string]interface{} {
	filters := make(map[string]interface{})

	filters["hospital_id"] = hospitalId
	queryParams := map[string]string{
		"first_name_th":  "first_name_th",
		"last_name_th":   "last_name_th",
		"email":          "email",
		"phone_number":   "phone_number",
		"first_name_en":  "first_name_en",
		"last_name_en":   "last_name_en",
		"middle_name_en": "middle_name_en",
		"middle_name_th": "middle_name_th",
		"patient_hn":     "patient_hn",
		"national_id":    "national_id",
		"passport_id":    "passport_id",
		"gender":         "gender",
		"date_of_birth":  "date_of_birth",
	}

	for dbField, queryParam := range queryParams {
		if value := c.Query(queryParam); value != "" {
			filters[dbField] = value
		}
	}

	return filters
}

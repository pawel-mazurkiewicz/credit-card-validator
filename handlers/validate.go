package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawel-mazurkiewicz/credit-card-validator/models"
	"github.com/pawel-mazurkiewicz/credit-card-validator/services"
)

func ValidateCardHandler(c *gin.Context) {
	var request models.ValidationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	isValid, cardType := services.ValidateCard(request.CardNumber)
	var message string
	if isValid {
		message = "Valid credit card number."
	} else {
		message = "Invalid credit card number."
	}

	response := models.ValidationResponse{
		IsValid:  isValid,
		CardType: cardType,
		Message:  message,
	}
	c.JSON(http.StatusOK, response)
}

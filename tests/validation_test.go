package tests

import (
	"testing"

	"github.com/pawel-mazurkiewicz/credit-card-validator/services"
)

func TestLuhnCheck(t *testing.T) {
	validCardNumbers := []string{
		"4111111111111111", // Visa
		"5555555555554444", // MasterCard
		"378282246310005",  // American Express
	}

	invalidCardNumbers := []string{
		"411111111111112",  // Invalid Visa
		"5500000000000005", // Invalid MasterCard
		"34000000000000",   // Invalid American Express
	}

	for _, number := range validCardNumbers {
		if isValid, cardType := services.ValidateCard(number); !isValid {
			t.Errorf("Expected card type %s number %s to be valid", cardType, number)
		}
	}

	for _, number := range invalidCardNumbers {
		if isValid, cardType := services.ValidateCard(number); isValid {
			t.Errorf("Expected card type %s number %s to be invalid", cardType, number)
		}
	}
}

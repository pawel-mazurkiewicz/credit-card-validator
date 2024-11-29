package tests

import (
	"testing"

	"github.com/pawel-mazurkiewicz/credit-card-validator/services"
)

func TestLuhnCheck(t *testing.T) {
	validCardNumbers := []string{
		"4111111111111111", // Visa
		"5500000000000004", // MasterCard
		"340000000000009",  // American Express
	}

	invalidCardNumbers := []string{
		"411111111111112",  // Invalid Visa
		"5500000000000005", // Invalid MasterCard
		"34000000000000",   // Invalid American Express
	}

	for _, number := range validCardNumbers {
		if isValid := services.ValidateCard(number); !isValid {
			t.Errorf("Expected card number %s to be valid", number)
		}
	}

	for _, number := range invalidCardNumbers {
		if isValid := services.ValidateCard(number); isValid {
			t.Errorf("Expected card number %s to be invalid", number)
		}
	}
}

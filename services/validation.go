package services

import (
    "regexp"
    "strings"
)

var cardPatterns = map[string]*regexp.Regexp{
    "Visa":        regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
    "MasterCard":  regexp.MustCompile(`^5[1-5][0-9]{14}$`),
    "American Express": regexp.MustCompile(`^3[47][0-9]{13}$`),
    // Add more card patterns as needed
}

// ValidateCard uses the Luhn algorithm to validate the card number
func ValidateCard(cardNumber string) (bool, string) {
    sanitizedNumber := sanitizeInput(cardNumber)
    isValid := luhnCheck(sanitizedNumber)
    cardType := detectCardType(sanitizedNumber)
    return isValid, cardType
}

func sanitizeInput(input string) string {
    return strings.ReplaceAll(strings.TrimSpace(input), " ", "")
}

func luhnCheck(cardNumber string) bool {
    sum := 0
    alternate := false
    for i := len(cardNumber) - 1; i >= 0; i-- {
        n := int(cardNumber[i] - '0')
        if alternate {
            n *= 2
            if n > 9 {
                n = (n % 10) + 1
            }
        }
        sum += n
        alternate = !alternate
    }
    return sum%10 == 0
}

func detectCardType(cardNumber string) string {
    for cardType, pattern := range cardPatterns {
        if pattern.MatchString(cardNumber) {
            return cardType
        }
    }
    return "Unknown"
}
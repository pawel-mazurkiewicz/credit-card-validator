package models

type ValidationResponse struct {
    IsValid  bool   `json:"isValid"`
    CardType string `json:"cardType,omitempty"`
    Message  string `json:"message"`
}
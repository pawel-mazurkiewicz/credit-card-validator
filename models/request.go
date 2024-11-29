package models

type ValidationRequest struct {
    CardNumber string `json:"cardNumber" binding:"required"`
}
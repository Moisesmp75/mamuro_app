package dto

import (
	"encoding/json"
	"errors"
	"time"
)

type EmailRequest struct {
	From    string    `json:"from" validate:"required"`
	To      []string  `json:"to" validate:"required"`
	Date    time.Time `json:"date" validate:"required"`
	Subject string    `json:"subject" validate:"required"`
	Content string    `json:"content" validate:"required"`
}

func CheckRequiredFields(req *EmailRequest) error {
	if req.From == "" {
		return errors.New("campo 'from' no puede estar vacío")
	}
	if len(req.To) == 0 {
		return errors.New("se requiere al menos un destinatario")
	}
	if req.Date.IsZero() {
		return errors.New("campo 'date' no puede estar vacío")
	}
	if req.Subject == "" {
		return errors.New("campo 'subject' no puede estar vacío")
	}
	if req.Content == "" {
		return errors.New("campo 'content' no puede estar vacío")
	}
	return nil
}

func ValidateEmailRequest(req *EmailRequest, body []byte) error {
	if err := json.Unmarshal(body, &req); err != nil {
		return err
	}
	if err := CheckRequiredFields(req); err != nil {
		return err
	}
	if err := validate.Struct(req); err != nil {
		return err
	}
	return nil
}

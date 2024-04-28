package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	RegisterRequest struct {
		Username  string     `json:"username" validate:"required"`
		Name      string     `json:"name" validate:"required"`
		Surname   string     `json:"surname" validate:"required"`
		Email     string     `json:"email" validate:"required"`
		Gender    string     `json:"gender" validate:"required"`
		BirthDate *time.Time `json:"birthDate" validate:"required"`
		Biography string     `json:"biography" validate:"required"`

		Password string `json:"password" validate:"required,min=8,max=20"`
	}

	RegisterResponse struct {
		Id          uuid.UUID  `json:"id"`
		CreatedDate *time.Time `json:"createdDate"`
	}
)

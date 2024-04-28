package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	GetProfileRequest struct{}

	GetProfileResponse struct {
		Id          uuid.UUID  `json:"id"`
		Username    string     `json:"username"`
		Name        string     `json:"name"`
		Surname     string     `json:"surname"`
		Email       string     `json:"email"`
		Gender      string     `json:"gender"`
		BirthDate   *time.Time `json:"birthDate"`
		CreatedDate *time.Time `json:"createdDate"`
		Biography   string     `json:"biography"`
	}
)

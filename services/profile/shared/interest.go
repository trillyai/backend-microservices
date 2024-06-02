package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateUserInterestRequest struct {
		Interests string `json:"interests"`
		Username  string
	}
	CreateUserInterestResponse struct {
		Interests []string `json:"interests"`
	}
)

type (
	DeleteUserInterestRequest struct {
		Interests string `json:"interests"`
		Username  string
	}
	DeleteUserInterestResponse struct {
		Interests []string `json:"interests"`
	}
)

type Interest struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	CreatedDate *time.Time `json:"createdDate"`
}

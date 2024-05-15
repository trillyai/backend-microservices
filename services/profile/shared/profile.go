package shared

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type (
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

type (
	UpdateProfileRequest struct {
		Username  string     `json:"username" validate:"required"`
		Name      string     `json:"name" validate:"required"`
		Surname   string     `json:"surname" validate:"required"`
		Email     string     `json:"email" validate:"required"`
		Gender    string     `json:"gender" validate:"required"`
		BirthDate *time.Time `json:"birthDate" validate:"required"`
		Biography string     `json:"biography" validate:"required"`
	}
	UpdateProfileResponse struct {
		Username       string     `json:"username"`
		LastUpdateDate *time.Time `json:"lastUpatedDate"`
	}
)

type (
	UploadProfileImageRequest struct {
		File *multipart.FileHeader `json:"file" validate:"required"`
	}
	UploadProfileImageResponse struct {
		Url string `json:"url"`
	}
)

package shared

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type (
	GetProfileResponse struct {
		Id           uuid.UUID  `json:"id"`
		Username     string     `json:"username"`
		Name         string     `json:"name"`
		Surname      string     `json:"surname"`
		Email        string     `json:"email"`
		Gender       string     `json:"gender"`
		CreatedDate  *time.Time `json:"createdDate"`
		Biography    string     `json:"biography"`
		ProfileImage string     `json:"profileImage,omitempty"`
	}
)

type (
	UpdateProfileRequest struct {
		Username     string
		Name         string `json:"name,omitempty"`
		Surname      string `json:"surname,omitempty"`
		Email        string `json:"email,omitempty"`
		Gender       string `json:"gender,omitempty"`
		Biography    string `json:"biography,omitempty"`
		ProfileImage string `json:"profileImage,omitempty"`
	}
	UpdateProfileResponse struct {
		Username        string     `json:"username"`
		LastUpdatedDate *time.Time `json:"lastUpatedDate"`
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

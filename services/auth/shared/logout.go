package shared

import "time"

type (
	LogoutRequest struct{}

	LogoutResponse struct {
		Username string     `json:"username"`
		EndDate  *time.Time `json:"endDate"`
	}
)

package shared

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}
)

package shared

type (
	CreateTripRequest struct {
		Areas    string `validate:"required"`
		Filters  string `validate:"required"`
		Distance string
	}

	CreateTripResponse struct {
		Root
	}
)

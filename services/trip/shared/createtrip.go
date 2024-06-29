package shared

type (
	CreateTripRequest struct {
		Areas    []string
		Filters  []string
		Distance float64
	}

	CreateTripResponse struct {
		Root
	}
)

package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/trip/contracts"
	"github.com/trillyai/backend-microservices/services/trip/shared"
)

type repository struct {
	logger logger.Logger
}

func NewRepository() contracts.Repository {
	return repository{
		logger: *logger.NewLogger("trip-repository"),
	}
}

// CreateTrip implements contracts.Repository.
func (r repository) CreateTrip(ctx context.Context, req shared.CreateTripRequest) (shared.CreateTripResponse, error) {
	claims := ctx.Value("user").(*auth.Claims)
	if claims.UserName == "" {
		return shared.CreateTripResponse{}, errors.New("context not found")
	}
	areas := strings.Split(req.Areas, ",")
	filters := strings.Split(req.Filters, ",")

	output, err := sendRequest(areas, filters, req.Distance)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateTripResponse{}, err
	}

	fmt.Printf("output: %v\n", output)

	return shared.CreateTripResponse{Root: output}, nil
}

// MakeRequest makes a GET request to the given URL with query parameters and returns the response body
func prepareRequest(baseURL string, params map[string]string) ([]byte, error) {
	// Parse the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// sendRequest constructs and sends a request with the given parameters and returns the parsed Root struct
func sendRequest(areas, filters []string, distance string) (shared.Root, error) {
	baseURL := "http://localhost:8000/api/generate"
	params := map[string]string{
		"filter":   join(filters, ","),
		"area":     join(areas, ","),
		"distance": distance,
	}

	body, err := prepareRequest(baseURL, params)
	if err != nil {
		return shared.Root{}, err
	}

	var root shared.Root
	err = json.Unmarshal(body, &root)
	if err != nil {
		return shared.Root{}, err
	}

	return root, nil
}

// join is a helper function to join a slice of strings with a given separator
func join(elems []string, sep string) string {
	result := ""
	for i, elem := range elems {
		if i > 0 {
			result += sep
		}
		result += elem
	}
	return result
}

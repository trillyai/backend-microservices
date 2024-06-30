package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/trillyai/backend-microservices/core/env"
	"github.com/trillyai/backend-microservices/services/trip/shared"
)

func prepareRequest(params map[string]string) ([]byte, error) {
	// Parse the base URL
	u, err := url.Parse(env.GenerateTripEndpoint)
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

func sendRequest(areas, filters []string, distance string) (shared.Root, error) {
	params := map[string]string{
		"filter":   join(filters, ","),
		"area":     join(areas, ","),
		"distance": distance,
	}

	body, err := prepareRequest(params)
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

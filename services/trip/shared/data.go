package shared

// Root structure
type Root struct {
	Data []Place `json:"data"`
	Info Info    `json:"info"`
}

// Place structure
type Place struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name,omitempty"`
	Lat        float64     `json:"lat"`
	Lon        float64     `json:"lon"`
	Filter     string      `json:"filter"`
	Area       string      `json:"area"`
	CloseNodes []CloseNode `json:"closeNodes,omitempty"`
}

// CloseNode structure
type CloseNode struct {
	ID     int64   `json:"id"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Filter string  `json:"filter"`
	Area   string  `json:"area"`
}

// Info structure
type Info map[string]AreaInfo

// AreaInfo structure
type AreaInfo struct {
	AvgLat float64        `json:"avgLat"`
	AvgLon float64        `json:"avgLon"`
	Counts map[string]int `json:"counts"`
}

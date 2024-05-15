package utils

import (
	"fmt"
	"strconv"
)

func GetOffSetAndLimit(offsetStr, limitStr string) (uint32, uint32, error) {
	offset, err := strconv.ParseUint(offsetStr, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid offset value: %s", err)
	}
	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid offset value: %s", err)
	}

	if limit > 250 {
		limit = 250
	}
	return uint32(offset), uint32(limit), nil
}

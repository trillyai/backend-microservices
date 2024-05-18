package utils

import (
	"fmt"
	"strconv"
)

func GetOffSetAndLimit(offsetStr, limitStr string) (uint32, uint32, error) {

	var offset, limit uint64
	if offsetStr == "" {
		offset = 0
	} else {
		os, err := strconv.ParseUint(offsetStr, 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid offset value: %s", err)
		}
		offset = os
	}

	if limitStr == "" {
		limit = 10
	} else {
		lm, err := strconv.ParseUint(limitStr, 10, 32)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid offset value: %s", err)
		}
		limit = lm
		if lm > 250 {
			limit = 250
		}
	}

	return uint32(offset), uint32(limit), nil

}

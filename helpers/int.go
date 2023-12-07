package helpers

import (
	"fmt"
	"strconv"
)

func ParseStringSliceToIntSlice(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))
	for i := range stringSlice {
		number, err := strconv.Atoi(stringSlice[i])
		if err != nil {
			return intSlice, fmt.Errorf("expected number from regex seed %s", stringSlice[i])
		}
		intSlice[i] = number
	}
	return intSlice, nil
}

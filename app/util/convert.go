package util

import (
	"fmt"
	"strconv"
)

func IntSliceToStringSlice(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))
	for i, v := range intSlice {
		stringSlice[i] = strconv.Itoa(v)
	}
	return stringSlice
}

func StringSliceToIntSlice(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))
	for i, v := range stringSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to int: %w", v, err)
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

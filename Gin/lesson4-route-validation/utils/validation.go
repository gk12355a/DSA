package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func ValidationStringLength(fieldName, value string, min, max int) error {
	length := len(value)
	if length < min || length > max {
		return fmt.Errorf("%s must be between %d and %d characters long", fieldName, min, max)
	}
	return nil
}
func ValidationRegex(fieldName, value string, re *regexp.Regexp, errorMessage string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s %s", errorMessage, fieldName)
	}
	return nil
}

func ValidationPositiveInt(fieldName, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer", fieldName)
	}
	if v <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", fieldName)
	}
	return v, nil

}

func ValidationUuid(fieldName, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", fieldName)
	}
	return uid, nil
}

func ValidationInList(fieldName, value string, allow map[string]bool) error {
	if !allow[value] {
		return fmt.Errorf("%s must be one of the following values: %v", fieldName, keys(allow))

	}
	return nil
}

func keys(m map[string]bool) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}

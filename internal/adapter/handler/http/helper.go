package http

import (
	"strconv"

	"github.com/sajadsalem/gostarter/internal/adapter/handler/http/response"
)

// stringToUint64 is a helper function to convert a string to uint64
func StringToUint64(str string) (uint64, error) {
	num, err := strconv.ParseUint(str, 10, 64)

	return num, err
}

// toMap is a helper function to add meta and data to a map
func ToMap(m response.Meta, data any, key string) map[string]any {
	return map[string]any{
		"meta": m,
		key:    data,
	}
}

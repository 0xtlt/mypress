package utils

import (
	"github.com/google/uuid"
)

// UUID function
func UUID() string {
	uuidWithHyphen := uuid.New()
	return uuidWithHyphen.String()
}

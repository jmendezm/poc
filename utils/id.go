package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

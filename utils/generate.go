package utils

import (
	"github.com/google/uuid"
)

//Generate TODO: complete id generation
func Generate() string {
	id := uuid.New().String()
	return id
}

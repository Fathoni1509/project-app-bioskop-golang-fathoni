package utils

import "github.com/google/uuid"

// GenerateToken membuat random string unik
func GenerateToken() string {
    return uuid.New().String()
}
package mapreduce

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length/2)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("could not generate random string: %w", err)
	}

	return hex.EncodeToString(randomBytes), nil
}

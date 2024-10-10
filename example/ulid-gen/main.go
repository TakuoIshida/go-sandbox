package main

import (
	"errors"
	"math/rand"

	"github.com/oklog/ulid/v2"

	"time"
)

var (
	t       = time.Unix(1000000, 0)
	entropy = ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
)

// generateUlid ex. output: 0000XSNJG0MQJHBF4QX1EFD6Y3
func generateUlid() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

// Get ex. output: 0000XSNJG0MQJHBF4QX1EFD6Y3
func Get() string {
	return generateUlid().String()
}

// Validation ulid string check size
func Validation(id string) error {
	if len(id) != ulid.EncodedSize {
		return errors.New("Unique ID generator validation error: length is not match")
	}
	return nil
}

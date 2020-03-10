package ec

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

// RandomResourceID generates a random string of 32 characters which emulates
// a real Elastic Cloud resource ID.
func RandomResourceID() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

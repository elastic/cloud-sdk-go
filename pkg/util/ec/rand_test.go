package ec

import "testing"

func TestRandomResourceID(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "is of 32 characters length",
			length: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomResourceID(); len(got) != tt.length {
				t.Errorf("RandomResourceID() = length %v, length is not %v", len(got), tt.length)
			}
		})
	}
}

package crypto

import (
	"reflect"
	"testing"
)

// TestHexToBytes tests the hexToBytes function.
func TestHexToBytes(t *testing.T) {
	tests := []struct {
		input    string
		expected []byte
		wantErr  bool
	}{
		{"4a6f686e", []byte{0x4a, 0x6f, 0x68, 0x6e}, false},
		{"abcdef", []byte{0xab, 0xcd, 0xef}, false},
		{"123456", []byte{0x12, 0x34, 0x56}, false},
		{"00ff", []byte{0x00, 0xff}, false},
		{"", []byte{}, false},
		{"4a6f686", nil, true},  // Odd length hex string
		{"4g6f686e", nil, true}, // Invalid hex character
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := hexToByte(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("hexToBytes() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

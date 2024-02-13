package protostream

import (
	"testing"
)

func TestDecodeVarint(t *testing.T) {
	tests := []struct {
		input    []byte
		expected uint64
		length   int
		name     string
	}{
		{[]byte{0x00}, 0, 1, "zero value"},
		{[]byte{0x01}, 1, 1, "positive value"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, length := DecodeVarint(test.input)
			if result != test.expected || length != test.length {
				t.Errorf("got result: %v, length: %v, expected result: %v, length: %v)", result, length, test.expected, test.length)
			}
		})
	}
}

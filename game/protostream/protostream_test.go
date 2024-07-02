// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

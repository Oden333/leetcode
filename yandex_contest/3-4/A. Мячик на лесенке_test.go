package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name:  "Пример 1",
			exp:   "7",
			input: `4`,
		},
		{
			name:  "Пример 2",
			exp:   "13",
			input: `5`,
		},
		{
			name:  "Пример 3",
			exp:   "1",
			input: `1`,
		},
		{
			name:  "Пример 4",
			exp:   "2",
			input: `2`,
		},
		{
			name:  "Пример 5",
			exp:   "4",
			input: `3`,
		},
		{
			name:  "Пример 6",
			exp:   "53798080",
			input: `30`,
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			A()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %q", tt.name, res)
				t.Fail()
			}
		})
	}
}

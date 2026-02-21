package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestE(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name:  "Пример 1",
			exp:   "44",
			input: "1 10",
		},
		{
			name:  "Пример 2",
			exp:   "10",
			input: "5 1",
		},
		{
			name:  "Пример 3",
			exp:   "46",
			input: "7 7",
		},
		{
			name:  "Пример 4",
			exp:   "46",
			input: "3 9",
		},
		{
			name:  "Пример 5",
			exp:   "1540",
			input: "1535 2346324236636443",
		},
		{
			name:  "Пример 6",
			exp:   "42",
			input: "4 7",
		},
		{
			name:  "Пример 7",
			exp:   "52",
			input: "7 8",
		},
		{
			name:  "Пример 8",
			exp:   "1",
			input: "1 0",
		},
		{
			name:  "Пример 9",
			exp:   "100",
			input: "100 124",
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			E()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

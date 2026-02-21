package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name:  "Пример 1",
			exp:   "5",
			input: `LLBLRRBRL`,
		},
		{
			name:  "Пример 2",
			exp:   "11",
			input: `BBBBBBBBBB`,
		},
		{
			name:  "Пример 3",
			exp:   "1",
			input: `RRR`,
		},
		{
			name:  "Пример 4",
			exp:   "1",
			input: `LLL`,
		},
		{
			name:  "Пример 5",
			exp:   "1",
			input: `R`,
		},
		{
			name:  "Пример 6",
			exp:   "2",
			input: `B`,
		},
		{
			name:  "Пример 7",
			exp:   "1",
			input: `LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL`,
		},
		{
			name:  "Пример 8",
			exp:   "1",
			input: `RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR`,
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			B()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %q", tt.name, res)
				t.Fail()
			}
		})
	}
}

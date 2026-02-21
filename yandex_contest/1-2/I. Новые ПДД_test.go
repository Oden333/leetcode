package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestI(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name: "Пример 1",
			exp:  "19",
			input: `8 7
3 4`,
		},
		{
			name: "Пример 2",
			exp:  "1",
			input: `3 4
4 3`,
		},
		{
			name: "Пример 3",
			exp:  "6",
			input: `8 7
11 7`,
		},
		{
			name: "Пример 4",
			exp:  "7",
			input: `8 7
10 9`,
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			I()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

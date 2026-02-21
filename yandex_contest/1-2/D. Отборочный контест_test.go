package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestD(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name:  "Пример 1",
			exp:   "15",
			input: "abacaba",
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			C()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_A(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name: "Пример 1",
			exp:  "1",
			input: `2
1 2`,
		},
		{
			name: "Пример 2",
			exp:  "2",
			input: `3
2 2 2`,
		},
		{
			name: "Пример 3",
			exp:  "10",
			input: `11
4 10 7 5 4 5 3 8 3 2 5`,
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
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

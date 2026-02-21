package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_B(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name:  "Пример 1",
			exp:   "0.50000",
			input: `1 2 2 10 10 10`,
		},
		{
			name:  "Пример 2",
			exp:   "1.20000",
			input: `4 1 2 5 5 5`,
		},
		{
			name:  "Пример 3",
			exp:   "1.49523",
			input: `2 3 4 7 6 5`,
		},
		{
			name:  "Пример 4",
			exp:   "1.271428571428571",
			input: `1 6 3 7 6 5`,
		},
		{
			name:  "Пример 5",
			exp:   "1.055555555555556",
			input: `2 3 4 10 9 2`,
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
			res := strings.TrimSpace(resWriter.String())
			if string([]byte(res[:6])) != string([]byte(tt.exp[:6])) {
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

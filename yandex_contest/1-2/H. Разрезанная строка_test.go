package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestH(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name: "Пример 1",
			exp:  "2 3 1",
			input: `12 3
cabacaqwerty
erty
caba
caqw
`,
		},
		{
			name: "Пример 2",
			exp:  "3 2 1",
			input: `3 3
aaa
a
a
a
  `,
		},
		{
			name: "Пример 3",
			exp:  "3 1 2",
			input: `12 3
cabacaqwerty
caqw
erty
caba
`,
		},
		{
			name: "Пример 4",
			exp:  "3 2 1",
			input: `12 3
cabacaqwerty
erty
caqw
caba
`,
		},
		{
			name: "Пример 5",
			exp:  "4 3 2 5 1",
			input: `15 5
abcabcabcqqqabc
abc
abc
abc
abc
qqq
`,
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			H()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %q", tt.name, res)
				t.Fail()
			}
		})
	}
}

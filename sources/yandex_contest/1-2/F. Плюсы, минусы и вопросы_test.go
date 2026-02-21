package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestF(t *testing.T) {
	tests := []struct {
		name, input, exp string
	}{
		{
			name: "Пример 1",
			exp:  "5",
			input: `4 3
+-+
??-
?-?
++?
`,
		},
		{
			name: "Пример 2",
			exp:  "12",
			input: `6 10
??+++-?-?-
-??+???--+
?-+?+-?+--
??????--?+
++--?--+-?
?-?+++?+-?
`,
		},
		{
			name: "Пример 3",
			exp:  "0",
			input: `1 1
+`,
		},
		{
			name: "Пример 4",
			exp:  "0",
			input: `1 1
?`,
		},
		{
			name: "Пример 5",
			exp:  "2",
			input: `2 2
?+
+?`,
		},
		{
			name: "Пример 6",
			exp:  "0",
			input: `2 2
++
++`,
		},
		{
			name: "Пример 7",
			exp:  "0",
			input: `2 2
--
--`,
		},
		{
			name: "Пример 8",
			exp:  "2",
			input: `1 3
?+?`,
		},
		{
			name: "Пример 9",
			exp:  "2",
			input: `3 1
?
?
?`,
		},
		{
			name: "Пример 10",
			exp:  "3",
			input: `2 3
???
???`,
		},
		{
			name: "Пример 11",
			exp:  "4",
			input: `3 3
+??
?+?
??+`,
		},
		{
			name: "Пример 12",
			exp:  "3",
			input: `3 2
??
??
??`,
		},
		{
			name: "Пример 13",
			exp:  "2",
			input: `2 2
??
??`,
		},
		{
			name: "Пример 14",
			exp:  "2",
			input: `2 2
?+
?+`,
		},
		{
			name: "Пример 15",
			exp:  "7",
			input: `6 3
++-
?-?
??-
+?+
---
---
`,
		},

		{
			name:  "Экстремальный случай - все вопросы",
			exp:   "1998", // n + m - 2? Нужно проверить
			input: "1000 1000\n" + strings.Repeat(strings.Repeat("?", 1000)+"\n", 1000),
		},
	}

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			F()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

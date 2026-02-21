package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestJ(t *testing.T) {

	var result []byte
	resWriter := bytes.NewBuffer(result)
	Output = resWriter

	for _, tt := range tests {
		Input = strings.NewReader(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			resWriter.Reset()
			J()
			if res := strings.TrimSpace(resWriter.String()); res != tt.exp {
				t.Logf("\nTest %q: \nOutput: %s", tt.name, res)
				t.Fail()
			}
		})
	}
}

var tests = []struct {
	name, input, exp string
}{
	{
		name: "Пример 1",
		exp:  "3",
		input: `3
List a = new List(2,3,5)
List b = a.subList(2,3)
b.get(1)`,
	},
	{
		name: "Пример 2",
		exp:  "16\n5",
		input: `5
List p = new List(2,4,8,16)
p.get(4)
List q = new List(3,9,27)
q.add(5)
q.get(4)
`,
	},
	{
		name: "Пример 3",
		exp:  "7\n2\n100\n100\n43\n14",
		input: `13
List x = new List(1,2,5,14,42)
List y = x.subList(1,4)
List z = y.subList(2,4)
y.set(1,7)
x.get(1)
z.get(1)
z.set(2,100)
x.get(3)
y.get(3)
x.add(132)
x.set(5,43)
x.get(5)
y.get(4)
`,
	},
	{
		name: "Пример 4",
		exp:  "1\n3\n10\n5",
		input: `6
List a = new List(1,2,3,4,5)
a.get(1)
a.get(3)
a.set(2,10)
a.get(2)
a.get(5)`,
	},
	{
		name: "Пример 5",
		exp:  "30\n20\n20\n30",
		input: `8
List base = new List(10,20,30,40,50,60)
List sub1 = base.subList(2,5)
List sub2 = sub1.subList(1,3)
base.get(3)
sub1.get(1)
sub2.get(1)
base.set(4,99)
sub2.get(2)`,
	},
	{
		name: "Пример 6",
		exp:  "1\n7\n9",
		input: `6
List lst = new List(1,3,5)
lst.add(7)
lst.add(9)
lst.get(1)
lst.get(4)
lst.get(5)`,
	},
	{
		name: "Пример 7",
		exp:  "50\n50\n10\n100",
		input: `10
List main = new List(2,4,6,8,10,12,14,16)
List view1 = main.subList(3,7)
List view2 = view1.subList(1,4)
List view3 = main.subList(5,8)
view1.set(2,50)
main.get(4)
view2.get(2)
view3.get(1)
view3.set(3,100)
main.get(7)`,
	},
}

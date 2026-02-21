package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/* var (
	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout
)

func main() {
	H()
} */

func H() {
	reader := bufio.NewReaderSize(Input, 1<<30)
	in, _ := reader.ReadString('\n')
	in = strings.Trim(in, "\n")

	nums := strings.Split(in, " ")
	len, _ := strconv.Atoi(nums[0])
	k, _ := strconv.Atoi(nums[1])

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")

	m := map[string][]int{}
	for i := range k {
		word, err := reader.ReadString('\n')
		if err != nil && str == "" {
			break
		}
		word = strings.Trim(word, "\n")

		m[word] = append(m[word], i+1)
	}

	delim := len / k
	for i := range k {
		key := str[delim*i : delim*(i+1)]
		nums := m[key]
		for i, n := range nums {
			if n == 0 {
				continue
			}

			writer.WriteString(fmt.Sprintf("%d", n))
			m[key][i] = 0
			break
		}

		if i+1 != k {
			writer.WriteString(" ")
		}
	}

	writer.WriteByte('\n')
}

/*

type z struct {
	order int
	idx   int
}

func H() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	ss, _ := reader.ReadString('\n')
	nums := strings.Split(strings.TrimSpace(ss), " ")
	_, _ = strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])

	s, _ := reader.ReadString('\n')
	splits := make([]z, 0, n2)
	for o := 1; o <= n2; o++ {
		s1, _ := reader.ReadString('\n')

		// splits = append(splits, s1)
		i := strings.Index(s, strings.TrimSpace(s1))
		// strings.Trim(s,)

		cur := z{
			idx:   i,
			order: o,
		}
		if len(splits) == 0 {
			splits = append(splits, cur)
			continue
		}
		for q, idz := range splits {
			if cur.idx < idz.idx {
				tmp := make([]z, len(splits)-q)
				copy(tmp, splits[q:])
				splits = append(append(splits[:q], cur), tmp...)
				break
			} else if q == len(splits)-1 {
				splits = append(splits, cur)
				break

			} else {
				continue
			}

		}
	}
	writer := bufio.NewWriterSize(Output, 1<<20)

	defer writer.Flush()
	for i, idz := range splits {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(fmt.Sprintf("%d", idz.order))
	}
	writer.WriteByte('\n')

}
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout
)

func main() {
	J()
}

// «new», «list», «set», «add» и «get»
func J() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	commCount, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(commCount))

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	lists := make(map[string][]string)
	subLists := make(map[string][]string)

	var getSubListValue func(subL string, idx int) string
	getSubListValue = func(subL string, idx int) string {
		if subSubL, ok := subLists[subL]; ok {
			offset, _ := strconv.Atoi(subSubL[0])
			offset--
			return getSubListValue(subSubL[2], idx+offset)
		}

		if List, ok := lists[subL]; ok {
			return List[idx]
		}
		panic("Zalupa")
	}

	var setSubListValue func(subL string, idx int, val string)
	setSubListValue = func(subL string, idx int, val string) {
		if subSubL, ok := subLists[subL]; ok {
			offset, _ := strconv.Atoi(subSubL[0])
			offset--
			setSubListValue(subSubL[2], idx+offset, val)
		}

		if List, ok := lists[subL]; ok {
			List[idx] = val
		}
	}

	for range n {
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if strings.HasPrefix(command, "List") {
			//List a = new List(2,3,5)
			//List b = a.subList(2,3)
			var (
				listName string
				listArgs string
			)
			comReader := strings.NewReader(command)
			switch {
			case strings.Contains(command, "new"):
				fmt.Fscanf(comReader, "List %s = new List(%s)", &listName, &listArgs)
				lists[listName] = strings.Split(strings.TrimSuffix(listArgs, ")"), ",")
			default:
				var (
					subCommand string
				)
				fmt.Fscanf(comReader, "List %s = %s", &listName, &subCommand)
				args := strings.Split(subCommand, ".")
				q := strings.TrimPrefix(args[1], "subList(")
				subLists[listName] = append(strings.Split(strings.TrimSuffix(q, ")"), ","), args[0])
			}
		} else {
			// y.get(3)
			// x.add(132)
			// x.set(5,43)

			actionReader := bufio.NewReader(strings.NewReader(command))
			listName, _ := actionReader.ReadString('.')
			listName = listName[:len(listName)-1]

			action, _ := actionReader.ReadString('(')
			switch strings.TrimSuffix(action, "(") {
			case "get":
				a, _ := actionReader.ReadString(')')
				arg, _ := strconv.Atoi(strings.TrimSuffix(a, ")"))
				arg--

				if subL, ok := subLists[listName]; ok {
					a, _ := strconv.Atoi(subL[0])
					a--

					// b, _ := strconv.Atoi(subL[1])
					// b--

					writer.WriteString(getSubListValue(subL[2], a+arg))
					writer.WriteByte('\n')
					continue
				}

				if L, ok := lists[listName]; ok {

					writer.WriteString(L[arg])
					writer.WriteByte('\n')
					continue
				}

			case "add":
				a, _ := actionReader.ReadString(')')
				if _, ok := lists[listName]; ok {
					lists[listName] = append(lists[listName], strings.TrimSuffix(a, ")"))
					continue
				}
			case "set":
				a, _ := actionReader.ReadString(')')
				args := strings.Split(strings.TrimSuffix(a, ")"), ",")

				idx, _ := strconv.Atoi(args[0])
				idx--
				val := args[1]

				if subL, ok := subLists[listName]; ok {
					a, _ := strconv.Atoi(subL[0])
					a--

					setSubListValue(subL[2], a+idx, val)
					continue
				}

				if L, ok := lists[listName]; ok {

					L[idx] = val
					continue
				}

			}
		}

	}
}

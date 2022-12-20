package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var stacks [][]string
	var stacksDone bool

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if stacks == nil {
			stacks = make([][]string, (len(line)/4)+1)
		}

		if !stacksDone {
			var j int
			var crate string
			for i := range stacks {
				j, crate = parseCrate(j, line)
				if crate != "" {
					stacks[i] = append(stacks[i], crate)
				}
				if j < 0 {
					stacksDone = true
					fileScanner.Scan() // scan empty line b/w moves
					break
				}
			}
		} else {
			n, x, y := parseMove(line)
			// 0 index
			x--
			y--

			for i := 0; i < n; i++ {
				// pop
				crate := stacks[x][0]
				stacks[x] = stacks[x][1:]

				// push
				stacks[y] = append([]string{crate}, stacks[y]...)
			}
		}
	}

	var tops string
	for _, stack := range stacks {
		tops += stack[0]
	}

	readFile.Close()
	fmt.Println(tops)
}

func parseMove(line string) (int, int, int) {
	parts := strings.Split(line, " ")

	n, _ := strconv.Atoi(parts[1])
	x, _ := strconv.Atoi(parts[3])
	y, _ := strconv.Atoi(parts[5])
	return n, x, y
}

func parseCrate(j int, line string) (int, string) {
	switch line[j] {
	case '\n':
	case '\t', ' ':
		if line[j+1] == '1' {
			return -1, ""
		}
		// empty crate
		j += 4
	case '[':
		char := line[j+1]
		j += 4
		return j, string(char)
	}
	return j, ""
}

func pop(stack []string) (string, []string) {
	return stack[0], stack[1:]
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

const ()

func f(in rune) int {
	if in < 'a' {
		return int(in - 'A' + 27)
	}
	return int(in - 'a' + 1)
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var priority int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		a := line[:len(line)/2]
		b := line[len(line)/2:]

		set := make(map[rune]bool, len(a))
		for _, char := range a {
			set[char] = false
		}
		for _, char := range b {
			if _, ok := set[char]; ok {
				set[char] = true
			}
		}
		for char, ok := range set {
			if ok {
				priority += f(char)
			}
		}
	}

	readFile.Close()
	fmt.Println(priority)
}

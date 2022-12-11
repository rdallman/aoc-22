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
		a := fileScanner.Text()
		fileScanner.Scan()
		b := fileScanner.Text()
		fileScanner.Scan()
		c := fileScanner.Text()

		ab := intersect(a, b)
		bc := intersect(ab, c)

		for _, char := range bc {
			priority += f(char)
		}
	}

	readFile.Close()
	fmt.Println(priority)
}

func intersect(a, b string) string {
	set := make(map[rune]bool, len(a))
	for _, char := range a {
		set[char] = false
	}

	for _, char := range b {
		if _, ok := set[char]; ok {
			set[char] = true
		}
	}

	var ret string
	for char, ok := range set {
		if ok {
			ret += string(char)
		}
	}
	return ret
}

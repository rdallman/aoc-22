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

	var p int
	for fileScanner.Scan() {
		line := fileScanner.Text()

		both := strings.Split(line, ",")

		a := strings.Split(both[0], "-")
		b := strings.Split(both[1], "-")

		aa, _ := strconv.Atoi(a[0])
		ab, _ := strconv.Atoi(a[1])
		ba, _ := strconv.Atoi(b[0])
		bb, _ := strconv.Atoi(b[1])

		if ab >= ba && bb >= aa {
			p++
		}
	}

	readFile.Close()
	fmt.Println(p)
}

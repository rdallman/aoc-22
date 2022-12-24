package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	buf := make([]string, 14)

	var i int
	for i = 0; fileScanner.Scan(); i++ {
		char := fileScanner.Text()

		if i < 14 {
			buf[i] = char
			if i < 13 {
				continue
			}
		} else {
			buf = buf[1:]
			buf = append(buf, char)
		}

		var repeat bool
		for j, b := range buf {
			// lazy n^2 check membership on every char
			for k := j + 1; k < len(buf); k++ {
				bb := buf[k]
				repeat = repeat || (bb == b && j != k)
			}
		}

		if !repeat {
			break
		}
	}

	readFile.Close()
	fmt.Println(i + 1)
}

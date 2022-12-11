package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var max, cur int
	for fileScanner.Scan() {
		num := fileScanner.Text()
		if num == "" {
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}

		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		cur += n
	}

	readFile.Close()

	fmt.Println(max)
}

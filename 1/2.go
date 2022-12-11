package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var cals []int
	var cur int
	for fileScanner.Scan() {
		num := fileScanner.Text()
		if num == "" {
			cals = append(cals, cur)
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

	sort.Ints(cals)
	cals = cals[len(cals)-3:]
	var sum int
	for _, c := range cals {
		sum += c
	}

	fmt.Println(sum)
}

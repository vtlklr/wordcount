package main

import (
	"bufio"
	"fmt"

	"os"
)

func main() {
	nums := readInput()

	fmt.Println(nums)
}

func readInput() int {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var a int
	for scanner.Scan() {

		a++
	}
	return a
}

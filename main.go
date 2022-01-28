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

/*func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.TrimSpace(s)
	s1 := strings.Split(s, " ")
	if len(s1) == 1 && s1[0] == "" {
		fmt.Print("0")
	} else {
		fmt.Print(len(s1))
	}
	//fmt.Println(len(s1))

}*/

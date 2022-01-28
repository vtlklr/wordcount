package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Trim(s, "\n")
	s1 := strings.Split(s, " ")
	fmt.Println(len(s1))
}

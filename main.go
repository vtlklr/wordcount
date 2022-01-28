package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	s = strings.TrimSpace(s)
	s1 := strings.Split(s, " ")
	fmt.Println(len(s1))
}

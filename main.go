package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	string_fields := strings.Fields(s)
	num := len(string_fields)
	ret := make(map[string]int)
	for i := 0; i < num; i++ {
		(ret[string_fields[i]])++
		fmt.Println(ret)
	}

	return ret
}

func main() {
	wc.Test(WordCount)
}

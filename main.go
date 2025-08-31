package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	text = Reverse(text)
	fmt.Println(text)
}

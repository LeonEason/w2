package main

import (
	"fmt"
	"sort"
)

type array struct {
	number int
}

type S []*array

func (s S) Len() int {
	return len(s)
}
func (s S) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s S) Less(i, j int) bool {
	return s[i].number < s[j].number
}

func main() {
	data := []*array{}
	var n int
	//flag := true
	for {
		fmt.Scanln(&n)
		if n == -1 {

			break
		}
		data = append(data, &array{n})
	}
	sort.Sort(S(data))
	for i := 0; i < len(data); i = i + 1 {
		fmt.Printf("%d", data[i])
	}
}

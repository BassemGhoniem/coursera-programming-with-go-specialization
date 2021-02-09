package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	for i := 0; true; i++ {
		fmt.Print("Please enter number [X to exit]: ")
		var s string
		fmt.Scanf("%s", &s)
		if s == "X" {
			break
		}
		number, _ := strconv.Atoi(s)
		if i < len(slice) {
			slice[i] = number
		} else {
			slice = append(slice, number)
		}
		sort.Ints(slice)
		fmt.Println(slice)
	}
}

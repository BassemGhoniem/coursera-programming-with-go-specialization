package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func trim(s string, length int) string {
	var size, x int

	for i := 0; i < length && x < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[x:])
		x += size
	}
	return s[:x]
}

func main() {
	type Person struct {
		fname string
		lname string
	}
	fmt.Print("Please enter file name: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	file, _ := os.Open(fileName)
	scanner = bufio.NewScanner(file)
	namesSlice := make([]Person, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(line, " ")
		namesSlice = append(namesSlice, Person{fname: trim(name[0], 20), lname: trim(name[1], 20)})
	}
	for _, p := range namesSlice {
		fmt.Println(p.fname, p.lname)
	}
	file.Close()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please enter string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.ToLower(scanner.Text())
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	person["name"] = scanner.Text()
	fmt.Print("Please enter your address: ")
	scanner.Scan()
	person["address"] = scanner.Text()
	barr, _ := json.Marshal(person)
	fmt.Println(barr)
	fmt.Println(string(barr))
}

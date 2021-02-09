package main

import "fmt"

func main() {
	fmt.Print("Please enter floating number: ")
	var x float32
	fmt.Scanf("%f", &x)
	fmt.Println("Your floating number truncated to:", int(x))
}

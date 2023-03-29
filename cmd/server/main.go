package main

import "fmt"

func Run() error {
	fmt.Println("Hello, World!")
	return nil
}

func main() {
	fmt.Println("Hello, World!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

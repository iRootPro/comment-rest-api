package main

import "fmt"

func Run() error {
	fmt.Println("Run services")
	return nil
}

func main() {
	fmt.Println("Start server")
	if err := Run(); err != nil {
		fmt.Println(err.Error())
	}
}

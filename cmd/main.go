package main

import "fmt"

// Run instantiate and start-up the application
func Run() error {
	fmt.Println("running....")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println("Error running the server :", err)
	}
}

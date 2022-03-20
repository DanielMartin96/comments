package main

import "fmt"

// Run - going to responsible for instantiation and start up of our application
func Run() error {
	fmt.Println("Starting our application")
	return nil
}

func main(){
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
package main

import (
	"context"
	"fmt"

	"github.com/DanielMartin96/comments/internal/db"
)

// Run - going to responsible for instantiation and start up of our application
func Run() error {
	fmt.Println("Starting our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully pinged database")
	return nil
}

func main(){
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
package main

import (
	"fmt"
	"log"
	"task-tracker/internal/commands"
)

func main() {
	cli := commands.New()
	result, err := cli.Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

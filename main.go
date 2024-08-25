package main

import (
	"fmt"
	"log"
	"task-tracker/cmd"
)

func main() {
	cli := cmd.New()
	result, err := cli.Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

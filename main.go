package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Let's Go Clone SQLite3 version 0.0.1\n")
	fmt.Print("Enter \".help\" for instructions\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("sqlite> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if input == ".help\n" {
			fmt.Print("noop\n")
		} else if input == ".exit\n" {
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command %s", input)
		}
	}
}

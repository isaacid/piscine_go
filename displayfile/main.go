package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	argsArr := os.Args
	argsArrLen := len(argsArr)

	if argsArrLen < 2 { // Handle the case where no file name is provided
		fmt.Println("File name missing")
		return
	}

	if argsArrLen > 2 { // Handle the case where too many arguments are provided
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(argsArr[1]) // Safely open the file, now that we know argsArr[1] exists
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file) // Read the file content
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf(string(content)) // Print the file content
}

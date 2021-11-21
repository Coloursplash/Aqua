package main


import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// get fileName from args
	argsLength := len(os.Args[1:])
	if (argsLength == 0) {
		printHelp()
		os.Exit(1)
	}

	fileName := os.Args[1]

	// read in file
	srcFile, err := os.Open(fileName)

	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	defer srcFile.Close()

	// read file to list
	scanner := bufio.NewScanner(srcFile)
	var code []string
    for scanner.Scan() {
        code = append(code, scanner.Text())
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
		os.Exit(1)
    }

	
}


func printHelp() {
	fmt.Println()
	fmt.Println("no you")
	fmt.Println()
}
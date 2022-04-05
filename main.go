package main

import "fmt"

var version string = "development"

func main() {
	printVersion()
	
}

func printVersion() {
	fmt.Printf("Starting go-api-sample %s\n\n", version)
}

package main

import (
		"fmt"
		"bufio"
		"os"
		"log"
		)

func main() {

	fmt.Println("Hi There")
	file, err := os.Open("test.tsv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
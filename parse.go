package main

import (
		"fmt"
		"bufio"
		"os"
		"log"
		"regexp"
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
		scanText := scanner.Text()
		regex := regexp.MustCompile("\t")
		splitted := regex.Split(scanText, -1)
		fmt.Println(splitted, len(splitted))
	}
}
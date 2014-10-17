package main

import (
		"fmt"
		"bufio"
		"os"
		"log"
		"regexp"
		"strings"
		)

func toCamelCase(str string) (string) {
	lower := strings.ToLower(str)
	regex := regexp.MustCompile("\\s+(\\w)")
	capit := regex.ReplaceAllStringFunc(lower, func(m string) (string) {
		return strings.TrimSpace(strings.ToUpper(m))
	})
	return capit
}

func main() {

	fmt.Println("Hi There")
	file, err := os.Open("test.tsv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	first := 1
	for scanner.Scan() {
		scanText := scanner.Text()
		regex := regexp.MustCompile("\t")
		splitted := regex.Split(scanText, -1)
		fmt.Println(splitted, len(splitted))
		if first == 1 {
			fmt.Println("first")
			first = 0
			for _, elem := range splitted {
				fmt.Println(elem)
				fmt.Println(toCamelCase(elem))
			}
		} else {
			fmt.Println("else")
		}
	}
}
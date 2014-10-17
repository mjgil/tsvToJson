package main

import (
		"fmt"
		"bufio"
		"os"
		"log"
		"regexp"
		"strings"
		"encoding/json"
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

	// fmt.Println("Hi There")
	file, err := os.Open("/Users/malcomgilbert/Downloads/Employees.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	first := 1
	firstRow := []string{}
	final := []map[string]string{}
	for scanner.Scan() {
		scanText := scanner.Text()
		regex := regexp.MustCompile("\t")
		splitted := regex.Split(scanText, -1)
		// fmt.Println(splitted, len(splitted))
		if first == 1 {
			firstRow = splitted
			first = 0
			for index, elem := range splitted {
				// fmt.Println(elem)
				// fmt.Println(toCamelCase(elem))
				firstRow[index] = toCamelCase(elem)
			}
		} else {
			m := make(map[string] string)
			// fmt.Println(firstRow)
			for index, elem := range splitted {
				// fmt.Println(elem)
				// fmt.Println(strings.Trim(elem, "\""))
				m[firstRow[index]] = strings.Trim(elem, "\"")
			}
			// fmt.Println(m)
			final = append(final, m)
			// fmt.Println(final)
		}
	}

	binaryJson, _ := json.MarshalIndent(final, "", "    ")
	fmt.Println(string(binaryJson))

	f, _ := os.Create("employees.json")
    defer f.Close()
    f.Write(binaryJson)
}
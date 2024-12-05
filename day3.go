package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sumOfValidExpressionsWithConditionals(str string) int {
	parts := strings.Split(str, "don't()")
	sum := 0
	for index, part := range parts {
		if index == 0 {
			sum += sumOfValidExpressionsInAString(part)
		} else if strings.Contains(part, "do()") {
			delimiter := "do()"
			doIndex := strings.Index(part, delimiter)
			strAfterDoStatement := part[doIndex+len(delimiter):]
			if len(strAfterDoStatement) > 1 {
				sum += sumOfValidExpressionsInAString(strAfterDoStatement)
			}
		}
	}
	return sum
}

func sumOfValidExpressionsInAString(str string) int {
	sum := 0
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}
	return sum
}

func day3() {
	fileName := "./inputs/day3.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	str := ""
	for fileScanner.Scan() {
		str += fileScanner.Text()
	}
	result := sumOfValidExpressionsWithConditionals(str)
	fmt.Println(result)
}

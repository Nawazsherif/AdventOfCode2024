package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countOccurrences(number int, numbers []int) int {
	count := 0
	for _, num := range numbers {
		if num == number {
			count++
		}
	}
	return count
}

func uniqueElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	unique := make([]int, 0)
	seen := make(map[int]bool)
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	return unique
}

func day1_part1() {
	fileName := "./inputs/day1.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	var list1 []int
	var list2 []int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		data := fileScanner.Text()
		parts := strings.Fields(data)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	var result float64
	for i, _ := range list1 {
		result += math.Abs(float64(list1[i] - list2[i]))
	}
	fmt.Print(result)
}

func day1_part2() {
	fileName := "./inputs/day1.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	var list1 []int
	var list2 []int

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		data := fileScanner.Text()
		parts := strings.Fields(data)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	list1 = uniqueElements(list1)

	var result float64
	for i := range list1 {
		result += float64(countOccurrences(list1[i], list2) * list1[i])
	}
	fmt.Print(result)
}

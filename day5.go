package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intersect(arr1, arr2 []int) []int {
	intersection := []int{}
	m := make(map[int]bool)

	for _, val := range arr1 {
		m[val] = true
	}

	for _, val := range arr2 {
		if m[val] {
			intersection = append(intersection, val)
		}
	}

	return intersection
}

func contains(arr []int, target int) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}

func stringArrayToIntArray(str []string) []int {
	var arr []int
	for _, s := range str {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}
	return arr
}

func appendOrAdd(pageRule map[int][]int, key int, value int) {
	if _, ok := pageRule[key]; ok {
		pageRule[key] = append(pageRule[key], value)
	} else {
		pageRule[key] = []int{value}
	}
}

func day5() {
	fileName := "./inputs/day5.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	beforePageRule := make(map[int][]int)
	afterPageRule := make(map[int][]int)
	var middlePageNumbers = 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if strings.Contains(text, "|") {
			updateRules(text, afterPageRule, beforePageRule)
		} else if strings.Contains(text, ",") {
			var pageOrder = stringArrayToIntArray(strings.Split(text, ","))
			isPageOrderCorrect := isPageOrderCorrect(pageOrder, afterPageRule, beforePageRule)
			middlePageNumbers = part1(isPageOrderCorrect, pageOrder, middlePageNumbers)
			//middlePageNumbers = part2(isPageOrderCorrect, pageOrder, afterPageRule, beforePageRule, middlePageNumbers)
		}
	}

	fmt.Println(middlePageNumbers)
}

func part2(isPageOrderCorrect bool, pageOrder []int, afterPageRule map[int][]int, beforePageRule map[int][]int, middlePageNumbers int) int {
	if !isPageOrderCorrect {
		sort.Slice(pageOrder, func(i, j int) bool {
			return contains(afterPageRule[pageOrder[i]], pageOrder[j])
		})
		sort.Slice(pageOrder, func(i, j int) bool {
			return contains(beforePageRule[pageOrder[j]], pageOrder[i])
		})
		mid := len(pageOrder) / 2
		pageNumber := pageOrder[mid]
		middlePageNumbers += pageNumber
	}
	return middlePageNumbers
}

func part1(isPageOrderCorrect bool, pageOrder []int, middlePageNumbers int) int {
	if isPageOrderCorrect {
		mid := len(pageOrder) / 2
		pageNumber := pageOrder[mid]
		middlePageNumbers += pageNumber
	}
	return middlePageNumbers
}

func isPageOrderCorrect(pageOrder []int, afterPageRule map[int][]int, beforePageRule map[int][]int) bool {
	var isPageOrderCorrect = true
	for index, page := range pageOrder {
		var pagesToBeAddedAfter = afterPageRule[page]
		var pagesAddedBeforeTheCurrentPage = pageOrder[:index]
		var pagesToBeAddedBefore = beforePageRule[page]
		var pagesAddedAfterTheCurrentPage = pageOrder[index+1:]
		pagesInWrongOrder :=
			append(intersect(pagesToBeAddedAfter, pagesAddedBeforeTheCurrentPage),
				intersect(pagesToBeAddedBefore, pagesAddedAfterTheCurrentPage)...)
		if len(pagesInWrongOrder) > 0 {
			isPageOrderCorrect = false
			break
		}
	}
	return isPageOrderCorrect
}

func updateRules(text string, afterPageRule map[int][]int, beforePageRule map[int][]int) {
	var str = strings.Split(text, "|")
	page1, _ := strconv.Atoi(str[0])
	page2, _ := strconv.Atoi(str[1])
	appendOrAdd(afterPageRule, page1, page2)
	appendOrAdd(beforePageRule, page2, page1)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isDescending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

func areLevelsSafeWithoutDampener(arr []int) bool {
	if isDescending(arr) || isAscending(arr) {
		for i := 1; i < len(arr); i++ {
			if math.Abs(float64(arr[i]-arr[i-1])) > 3 || math.Abs(float64(arr[i]-arr[i-1])) < 1 {
				return false
			}
		}
		return true
	}
	return false
}

func areLevelsSafeWithDampener(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		updatedLevel := append([]int{}, arr[:i]...)
		updatedLevel = append(updatedLevel, arr[i+1:]...)

		if isDescending(updatedLevel) || isAscending(updatedLevel) {
			var isUpdatedLevelSafe = true
			for j := 1; j < len(updatedLevel); j++ {
				if math.Abs(float64(updatedLevel[j]-updatedLevel[j-1])) > 3 || math.Abs(float64(updatedLevel[j]-updatedLevel[j-1])) < 1 {
					isUpdatedLevelSafe = false
				}
			}
			if isUpdatedLevelSafe {
				return true
			}
		}
	}
	return false
}

func isAscending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func stringToIntArray(str string) []int {
	var arr []int
	for _, s := range strings.Fields(str) {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}
	return arr
}

func day2() {
	fileName := "./inputs/day2.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	var countMap = make(map[string]int)
	countMap["safe"] = 0
	countMap["un_safe"] = 0

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		data := fileScanner.Text()
		var list = stringToIntArray(data)
		if areLevelsSafeWithDampener(list) {
			countMap["safe"] += 1
		} else {
			countMap["un_safe"] += 1
		}
	}

	fmt.Print(countMap["safe"])
}

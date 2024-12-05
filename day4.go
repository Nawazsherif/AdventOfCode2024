package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkVertically(list []string, rowIndex int, colIndex int) bool {
	rows := len(list)
	if rowIndex < (rows - 3) {
		str := string(list[rowIndex][colIndex])
		for i := 1; i < 4; i++ {
			str += string(list[rowIndex+i][colIndex])
		}
		return isEqualToXMASOrSAMX(str)
	}
	return false
}

func checkHorizontally(list []string, row int, col int) bool {
	cols := len(list[0])
	if col < (cols - 3) {
		str := string(list[row][col])
		for i := 1; i < 4; i++ {
			str += string(list[row][col+i])
		}
		return isEqualToXMASOrSAMX(str)
	}
	return false
}

func checkRightDiagonal(list []string, row int, col int) bool {
	cols := len(list[0])
	rows := len(list)
	if (col < (cols - 3)) && (row < (rows - 3)) {
		str := string(list[row][col])
		for i := 1; i < 4; i++ {
			str += string(list[row+i][col+i])
		}
		return isEqualToXMASOrSAMX(str)
	}
	return false
}

func checkLeftDiagonal(list []string, row int, col int) bool {
	rows := len(list)
	if (col > 2) && (row < (rows - 3)) {
		str := string(list[row][col])
		for i := 1; i < 4; i++ {
			str += string(list[row+i][col-i])
		}
		return isEqualToXMASOrSAMX(str)
	}
	return false
}

func isEqualToXMASOrSAMX(str string) bool {
	return str == "XMAS" || str == "SAMX"
}

func isEqualToMASOrSAM(str string) bool {
	return str == "MAS" || str == "SAM"
}

func day4_part1() {
	fileName := "./inputs/day4.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	var list []string
	for fileScanner.Scan() {
		list = append(list, fileScanner.Text())
	}
	count := 0
	for rowIndex, _ := range list {
		for colIndex := 0; colIndex < len(list[0]); colIndex++ {
			char := string(list[rowIndex][colIndex])
			if (char == "X") || (char == "S") {
				if checkVertically(list, rowIndex, colIndex) {
					count++
				}
				if checkHorizontally(list, rowIndex, colIndex) {
					count++
				}
				if checkRightDiagonal(list, rowIndex, colIndex) {
					count++
				}
				if checkLeftDiagonal(list, rowIndex, colIndex) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func isXMASFound(list []string, rowIndex int, colIndex int) bool {
	rows := len(list)
	var count = 0
	cols := len(list[0])
	// check right diagonal
	if (colIndex < (cols - 2)) && (rowIndex < (rows - 2)) {
		str := string(list[rowIndex][colIndex])
		for i := 1; i < 3; i++ {
			str += string(list[rowIndex+i][colIndex+i])
		}
		if isEqualToMASOrSAM(str) {
			count++
		}
	}
	// check left diagonal
	newColIndex := colIndex + 2
	if (newColIndex > 1) && (newColIndex < cols) && (rowIndex < (rows - 2)) {
		str := string(list[rowIndex][newColIndex])
		for i := 1; i < 3; i++ {
			str += string(list[rowIndex+i][newColIndex-i])
		}
		if isEqualToMASOrSAM(str) {
			count++
		}
	}
	return count == 2
}

func day4_part2() {
	fileName := "./inputs/day4.txt"
	file, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	var list []string
	for fileScanner.Scan() {
		list = append(list, fileScanner.Text())
	}
	count := 0
	for rowIndex, _ := range list {
		for colIndex := 0; colIndex < len(list[0]); colIndex++ {
			char := string(list[rowIndex][colIndex])
			if char == "M" || char == "S" {
				if isXMASFound(list, rowIndex, colIndex) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

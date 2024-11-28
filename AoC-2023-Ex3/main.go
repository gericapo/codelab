package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var schematic [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := len(schematic)
	cols := len(schematic[0])
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var totalSum int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if unicode.IsDigit(schematic[r][c]) {
				num, _ := strconv.Atoi(string(schematic[r][c]))
				isPartNumber := false

				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						if !unicode.IsDigit(schematic[nr][nc]) && schematic[nr][nc] != '.' {
							isPartNumber = true
							break
						}
					}
				}

				if isPartNumber {
					totalSum += num
				}
			}
		}
	}

	fmt.Println("Sum of all part numbers:", totalSum)
}


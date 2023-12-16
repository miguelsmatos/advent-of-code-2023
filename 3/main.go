package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Part struct {
	value int
	row   int
	i     int
	j     int
}

type Gear struct {
	row    int
	column int
}

func hasSymbolInRange(str string, i, j int) bool {
	if len(str) == 0 {
		return false
	}
	ii := i
	if i < 0 {
		ii = 0
	}
	jj := j
	if j > len(str) {
		jj = len(str)
	}
	for index := ii; index < jj; index++ {
		if str[index] != '.' && (str[index] < '0' || str[index] > '9') {
			return true
		}
	}
	return false
}

func main() {

	filename := "input.txt"

	file, error := os.Open(filename)

	if error != nil {
		fmt.Printf("File %v could not be opened", filename)
		return
	}

	defer file.Close()

	line_count := 0

	scanner := bufio.NewScanner(file)

	sum := 0

	var parts []Part
	var gears []Gear

	scanner.Scan()
	line := ""
	next_line := scanner.Text()
	end := false

	for !end {
		previous_line := line
		line = next_line
		if scanner.Scan() {
			next_line = scanner.Text()
		} else {
			end = true
			next_line = ""
		}
		line_length := len(line)
		reading_number := false
		i := 0
		for j := 0; j < line_length; j++ {
			if line[j] == '*' {
				g := Gear{line_count, j}
				gears = append(gears, g)
			}
			if !unicode.IsDigit(rune(line[j])) {
				if reading_number {
					value, _ := strconv.Atoi(line[i:j])
					if hasSymbolInRange(previous_line, i-1, j+1) || hasSymbolInRange(next_line, i-1, j+1) || hasSymbolInRange(line, i-1, j+1) {
						sum += value
						p := Part{value, line_count, i, j - 1}
						parts = append(parts, p)
					}
					reading_number = false
				}
			} else if j == line_length-1 {
				if reading_number {
					value, _ := strconv.Atoi(line[i:line_length])
					if hasSymbolInRange(previous_line, i-1, line_length) || hasSymbolInRange(next_line, i-1, line_length) || hasSymbolInRange(line, i-1, line_length) {
						sum += value
						p := Part{value, line_count, i, line_length - 2}
						parts = append(parts, p)
					}
				}
				reading_number = false

			} else {
				if !reading_number {
					i = j
					reading_number = true
				}
			}
		}
		line_count++
		fmt.Printf("Line %03d: sum %d\n", line_count, sum)
	}

	gear_ratio_sum := 0
	for g, gear := range gears {
		gear_ratio := 1
		parts_for_gear := 0
		fmt.Printf("-> Gear %d [%d.%d]\n", g, gear.row, gear.column)
		for p, part := range parts {
			if part.row < gear.row-1 {
				continue
			}
			if part.row > gear.row+1 {
				break
			}
			if gear.column >= part.i-1 && gear.column <= part.j+1 {
				fmt.Printf("  * Part %d %d:[%d-%d] (%d)\n", p, part.row, part.i, part.j, part.value)
				parts_for_gear += 1
				gear_ratio *= part.value
			}
			if parts_for_gear > 2 {
				break
			}
		}
		if parts_for_gear != 2 {
			gear_ratio = 0
		}
		fmt.Printf("   => %d parts: ratio: %v\n", parts_for_gear, gear_ratio)

		gear_ratio_sum += gear_ratio
	}
	fmt.Printf("=> Total gear ratio: %v\n", gear_ratio_sum)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func substringToInt(str string, begin bool) int {
	if len(str) < 3 {
		return -1
	}
	relevantFunc := strings.HasPrefix
	if !begin {
		relevantFunc = strings.HasSuffix
	}

	if len(str) >= 3 {
		if relevantFunc(str, "one") {
			return 1
		}
		if relevantFunc(str, "two") {
			return 2
		}
		if relevantFunc(str, "six") {
			return 6
		}
	}
	if len(str) >= 4 {
		if relevantFunc(str, "four") {
			return 4
		}
		if relevantFunc(str, "five") {
			return 5
		}
		if relevantFunc(str, "nine") {
			return 9
		}
	}
	if len(str) >= 5 {
		if relevantFunc(str, "three") {
			return 3
		}
		if relevantFunc(str, "seven") {
			return 7
		}
		if relevantFunc(str, "eight") {
			return 8
		}
	}
	return -1
}

func main() {

	filename := "input.txt"

	file, error := os.Open(filename)

	if error != nil {
		fmt.Printf("File %v could not be opened", filename)
		return
	}

	defer file.Close()

	first := 0
	second := 0
	line_count := 0
	sum := 0
	line_value := 0
	var aux int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		start_ptr := 0
		end_ptr := 0
		line_length := len(line)
		first = 0
		second = 0
		var r rune

		for start_ptr < line_length {
			r = rune(line[start_ptr])
			if unicode.IsDigit(rune(line[start_ptr])) {
				first = int(r - '0')
				break
			} else if start_ptr > 0 {
				aux = substringToInt(line[0:start_ptr+1], false)
				if aux > 0 {
					first = aux
					break
				}
			}
			start_ptr++
		}

		for end_ptr < line_length {
			r := rune(line[line_length-end_ptr-1])
			if unicode.IsDigit(r) {
				second = int(r - '0')
				break
			} else if end_ptr > 0 {
				aux = substringToInt(line[line_length-end_ptr-1:line_length], true)
				if aux > 0 {
					second = aux
					break
				}
			}
			end_ptr++
		}
		line_value = 10*first + second
		fmt.Printf("#%03d: first: %d last %d (%d)\n", line_count, first, second, line_value)
		sum += line_value
		line_count++
	}

	fmt.Printf("=> The total sum is %d⚡\n", sum)

}

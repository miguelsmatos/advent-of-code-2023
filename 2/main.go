package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var colorRE = regexp.MustCompile(`((?P<green>\d+) green)|((?P<blue>\d+) blue)|((?P<red>\d+) red)`)

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
	var rs string
	red, blue, green := 0, 0, 0
	round_red, round_blue, round_green := 0, 0, 0
	total_red, total_green, total_blue := 12, 13, 14
	max_red, max_green, max_blue := 0, 0, 0
	round_possible := true
	possible_rounds_sum := 0
	sum_of_powers := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_games := strings.Split(line, ": ")
		if len(line_games) != 2 {
			continue
		}
		game_number := line_count + 1
		rounds := strings.Split(line_games[1], ";")
		round_red, round_blue, round_green = 0, 0, 0
		max_red, max_green, max_blue = 0, 0, 0
		round_possible = true
		for _, round := range rounds {
			red, blue, green = 0, 0, 0
			matches := colorRE.FindAllStringSubmatch(round, 3)
			for _, match := range matches {
				rs = match[colorRE.SubexpIndex("red")]
				if len(rs) > 0 {
					intV, _ := strconv.Atoi(rs)
					red = intV
				}
				rs = match[colorRE.SubexpIndex("green")]
				if len(rs) > 0 {
					intV, _ := strconv.Atoi(rs)
					green = intV
				}
				rs = match[colorRE.SubexpIndex("blue")]
				if len(rs) > 0 {
					intV, _ := strconv.Atoi(rs)
					blue = intV
				}
			}
			round_red += red
			round_green += green
			round_blue += blue
			if red > total_red || green > total_green || blue > total_blue {
				round_possible = false
			}
			if red > max_red {
				max_red = red
			}
			if green > max_green {
				max_green = green
			}
			if blue > max_blue {
				max_blue = blue
			}
		}
		sum_of_powers += max_red * max_blue * max_green
		fmt.Printf("❓ ~~> Game %d is possible with %d reds, %d greens %d blues\n", game_number, max_red, max_green, max_blue)

		if round_possible {
			fmt.Printf("✅ ~~> Game %d is possible\n", game_number)
			possible_rounds_sum += game_number
		} else {
			fmt.Printf("❌ ~~> Game %d is impossible\n", game_number)
		}

		line_count++
	}

	fmt.Printf("=> Sum of possible games: \n   %d\n⚾🏉\n", possible_rounds_sum)
	fmt.Printf("=> Sum of powers: \n   %d\n⚾🏉\n", sum_of_powers)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	min := 0
	red := 12
	green := 13
	blue := 14

	f, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		isOK := true
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		line := scanner.Text()
		split := strings.Split(line, ":")
		gameID := strings.Split(split[0], " ")
		results := strings.Split(split[1], ";")
		for _, result := range results {
			count := strings.Split(result, ",")
			for _, values := range count {
				value := strings.Split(values, " ")
				switch value[2] {
				case "red":
					v, _ := strconv.Atoi(value[1])
					if v > maxRed {
						maxRed = v
					}
					if v > red {
						isOK = false
					}
				case "blue":
					v, _ := strconv.Atoi(value[1])
					if v > maxBlue {
						maxBlue = v
					}
					if v > blue {
						isOK = false
					}
				case "green":
					v, _ := strconv.Atoi(value[1])
					if v > maxGreen {
						maxGreen = v
					}
					if v > green {
						isOK = false
					}
				}
			}

		}

		power := maxGreen * maxBlue * maxRed
		min += power

		if isOK {
			v, _ := strconv.Atoi(gameID[1])
			sum += v
		}
	}
	fmt.Printf("Sum = %d\n", sum)
	fmt.Printf("Min = %d\n", min)
}

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
					if v > red {
						isOK = false
					}
				case "blue":
					v, _ := strconv.Atoi(value[1])
					if v > blue {
						isOK = false
					}
				case "green":
					v, _ := strconv.Atoi(value[1])
					if v > green {
						isOK = false
					}
				}
			}

		}

		if isOK {
			v, _ := strconv.Atoi(gameID[1])
			sum += v
		}
	}
	fmt.Printf("Sum = %d\n", sum)
}

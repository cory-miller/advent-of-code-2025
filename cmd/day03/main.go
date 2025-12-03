package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func partOne(battery string) int {
	maxDig := 0
	maxDigIdx := 0

	for i := 0; i < len(battery)-1; i++ {
		if int(battery[i]-'0') > maxDig {
			maxDig = int(battery[i] - '0')
			maxDigIdx = i
		}

		if maxDig == 9 {
			break
		}
	}

	minDig := 0

	for i := len(battery) - 1; i > maxDigIdx; i-- {
		if int(battery[i]-'0') > minDig {
			minDig = int(battery[i] - '0')
		}

		if minDig == 9 {
			break
		}
	}

	return maxDig*10 + minDig
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/03.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	partOneSum := 0

	for scanner.Scan() {
		battery := scanner.Text()

		partOneSum += partOne(battery)
	}

	fmt.Printf("Part One: %d\n", partOneSum)
}

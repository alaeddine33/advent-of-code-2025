package adventofcode2025

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day3() {
	fmt.Println("----------------DAY 3-----------------")

	file, err := os.Open("./inputs/day3.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltage := 0

	for scanner.Scan() {
		bank := strings.Split(scanner.Text(), "")

		largest := 0

		for i := range bank {

			currentLargest := 0

			current, _ := strconv.Atoi(bank[i])

			for j := i + 1; j < len(bank); j++ {

				next, _ := strconv.Atoi(bank[j])

				newLargest := current*10 + next

				if newLargest > currentLargest {
					currentLargest = newLargest
				}
			}
			if currentLargest > largest {
				largest = currentLargest
			}
		}

		fmt.Println(largest)
		joltage += largest
	}

	fmt.Printf("Total joltage = %d", joltage)
}

// func remove(name string, bank []string) []string {
// 	i := 0
// 	for idx, item := range bank {
// 		if item != name {
// 			bank[i] = bank[idx]
// 			i++
// 		}
// 	}
// 	return bank[:i]
// }

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

		largestBattery, _ := strconv.Atoi(bank[0])

		for i := range bank {

			currentLargest := 0

			current, _ := strconv.Atoi(bank[i])

			if i > 0 && largestBattery >= current {
				continue
			}

			largestBattery = current

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

func Day3Part2() {
	fmt.Println("----------------DAY 3 PART 2-----------------")

	file, err := os.Open("./inputs/day3.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	const NEED = 12

	for scanner.Scan() {
		bank := strings.Split(scanner.Text(), "")
		n := len(bank)

		largestSeq := ""

		for i := 0; i <= n-NEED; i++ {

			seq := bank[i]
			remainingNeeded := NEED - 1
			nextPos := i + 1
			for remainingNeeded > 0 {

				bestDigit := "-1"
				bestIndex := -1

				limit := n - remainingNeeded

				for j := nextPos; j <= limit; j++ {
					if bank[j] > bestDigit {
						bestDigit = bank[j]
						bestIndex = j
					}
				}

				seq += bestDigit
				nextPos = bestIndex + 1
				remainingNeeded--
			}

			if seq > largestSeq {
				largestSeq = seq
			}
		}

		fmt.Println(largestSeq)

		val, _ := strconv.Atoi(largestSeq)
		total += val
	}

	fmt.Printf("Total joltage (Part 2) = %d\n", total)
}

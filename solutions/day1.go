package adventofcode2025

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day1() {

	fmt.Println("----------------DAY 1-----------------")

	var password, dial int = 0, 50

	file, err := os.Open("./inputs/day1.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		number, err := strconv.Atoi((scanner.Text()[1:]))
		originalDial := dial

		if number >= 100 {
			password += number / 100
			number = number % 100
		}

		if err != nil {
			log.Fatalf("error %s", err)
		}

		if strings.HasPrefix(scanner.Text(), "L") {

			dial -= number

			if dial < 0 {
				dial += 100
				if dial > 0 && originalDial != 0 {
					password++
				}
			}

		}

		if strings.HasPrefix(scanner.Text(), "R") {

			dial += number

			if dial > 99 {
				dial -= 100
				if dial > 0 {
					password++
				}
			}
		}

		if dial == 0 {
			password++
		}
	}

	fmt.Println("Password", password)

}

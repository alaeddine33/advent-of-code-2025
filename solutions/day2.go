package adventofcode2025

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	fmt.Println("----------------DAY 2-----------------")

	file, err := os.Open("./inputs/day2.csv")
	if err != nil {
		log.Fatalf("error %s", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	ranges, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error %s", err)
		return
	}

	var invalidIDs []int

	for _, r := range ranges[0] {

		ids := strings.Split(r, "-")

		leftNumber, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		rightNumber, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		for i := leftNumber; i <= rightNumber; i++ {

			stringNumber := strconv.Itoa(i)
			numberLength := len(stringNumber)
			factors := getFactors(numberLength)

			for _, f := range factors {
				//comment for part 2
				// if numberLength%2 != 0 {
				// 	continue
				// }
				if f == numberLength {
					continue
				}
				if numberLength%f != 0 {
					continue
				}

				part := stringNumber[:f]
				repeatCount := numberLength / f
				//comment for part 2

				// if repeatCount%2 != 0 {
				// 	continue
				// }
				newNumber := strings.Repeat(part, repeatCount)

				if newNumber == stringNumber {
					invalidIDs = append(invalidIDs, i)
					break
				}
			}
		}

	}

	var sum int = 0
	for _, id := range invalidIDs {
		fmt.Println(id)
		sum += id
	}

	fmt.Println("Sum", sum)

}

func getFactors(n int) []int {
	var factors []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}
	return factors
}

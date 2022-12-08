package challenges

import (
	"strconv"
	"strings"
)

func Day01Part1(input string) (string, error) {
	elves := strings.Split(input, "\n\n")

	maxCalories := 0

	for _, elf := range elves {
		totalCalories := 0
		for _, caloriesStr := range strings.Split(elf, "\n") {
			val, err := strconv.Atoi(caloriesStr)
			if err != nil {
				return "", err
			}
			totalCalories += val
		}

		if totalCalories > maxCalories {
			maxCalories = totalCalories
		}
	}

	return strconv.Itoa(maxCalories), nil
}

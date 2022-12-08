package challenges

import (
	"sort"
	"strconv"
	"strings"
)

func Day01Part2(input string) (string, error) {
	elves := strings.Split(input, "\n\n")

	caloriesByElf := make([]int, 0)

	for _, elf := range elves {
		totalCalories := 0
		for _, caloriesStr := range strings.Split(elf, "\n") {
			val, err := strconv.Atoi(caloriesStr)
			if err != nil {
				return "", err
			}
			totalCalories += val
		}

		caloriesByElf = append(caloriesByElf, totalCalories)
	}

	sort.Ints(caloriesByElf)

	sum := 0

	for _, calories := range caloriesByElf[len(caloriesByElf)-3:] {
		sum += calories
	}

	return strconv.Itoa(sum), nil
}

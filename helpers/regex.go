package helpers

import (
	"regexp"
	"strconv"
)

var numGroupRegex = regexp.MustCompile(`(\d+)`)

func GetNumbersFromString(input string) ([]int, error) {
	nums := numGroupRegex.FindAllString(input, -1)
	var numbers = make([]int, len(nums))

	for i := range nums {
		number, err := strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}

func GetStringNumbersFromString(input string) []string {
	return numGroupRegex.FindAllString(input, -1)
}

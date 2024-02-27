package types

import (
	"fmt"
	"strconv"
)

func Average() float64 {

	var sum float64
	var count int

	for {
		var input string
		var value float64

		fmt.Print("Enter a value: ")
		fmt.Scanln(&input)

		if input == "" {
			break
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Print("Invalid input, try again: ")
			continue
		}

		sum += value
		count++
	}

	if count > 0 {
		return sum / float64(count)
	}

	return 0.0
}

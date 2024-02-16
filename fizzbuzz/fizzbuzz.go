package fizzbuzz

import "strconv"

func FizzBuzz(input []int) string {
	if len(input) == 0 {
		return "Empty Array"
	}

	message := ""

	for index, number := range input {
		if number <= 0 {
			return "Error"
		}

		comma := ","

		if index == len(input)-1 {
			comma = ""
		}

		if number%3 == 0 || number%5 == 0 {
			if number%3 == 0 {
				message += "Fizz"
			}

			if number%5 == 0 {
				message += "Buzz"
			}
		} else {
			message += strconv.Itoa(number)
		}

		message += comma
	}

	return message
}

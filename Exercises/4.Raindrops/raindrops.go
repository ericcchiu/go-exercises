package raindrops

import "strconv"

// Convert function that takes an integer and returns Pling, Plang, Plong if number are divisble by 3,5,7 and returns stringified number if not.
func Convert(num int) (output string) {

	if num%3 == 0 {
		output += "Pling"
	}

	if num%5 == 0 {
		output += "Plang"
	}

	if num%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		output = strconv.Itoa(num)
	}

	return
}

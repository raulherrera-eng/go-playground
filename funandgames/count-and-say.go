package funandgames

import "strconv"

func CountAndSay(n int) string {
	if n <= 0 {
		return "" // error ??
	}

	result := "1"
	for i := 1; i < n; i++ {
		result = generateNextSequence(result)
	}

	return result
}

func generateNextSequence(previousSequence string) string {
	var result []rune
	var counter = 0
	for i := 0; i < len(previousSequence); i++ {
		current := previousSequence[i]
		if (i+1) < len(previousSequence) && current == previousSequence[i+1] { // if next digit is the same
			counter++
			continue
		}

		result = append(result, []rune(strconv.Itoa(counter+1))...)
		result = append(result, rune(current))
		counter = 0
	}

	return string(result)
}

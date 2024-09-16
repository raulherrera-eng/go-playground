package funandgames

func FindLongestPalindrome(input string) string {

	if len(input) == 0 {
		return ""
	} else if len(input) == 1 {
		return input
	} else if len(input) == 2 {
		if input[0] == input[1] {
			return input
		} else {
			return string(input[0])
		}
	}

	var longestPalindrome string = ""
	for i := 0; i < len(input); i++ {

		if len(longestPalindrome) >= len(input)-i {
			return longestPalindrome
		}

		longestPalindrome = findLongestPalindromeAroundCenter(input, longestPalindrome, i)
	}

	return longestPalindrome
}

func findLongestPalindromeAroundCenter(input, longestPalindrome string, center int) string {

	// odd-length palindromes (centered at center)
	oddPalindrome := expandAroundCenter(input, center, center)
	if len(oddPalindrome) > len(longestPalindrome) {
		longestPalindrome = oddPalindrome
	}

	// even-length palindromes (centered between center and center+1)
	evenPalindrome := expandAroundCenter(input, center, center+1)
	if len(evenPalindrome) > len(longestPalindrome) {
		longestPalindrome = evenPalindrome
	}

	return longestPalindrome

}

// Helper function to expand around center and find the longest palindrome
func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

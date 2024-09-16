package funandgames

import "testing"

func TestFindLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"},               // case with multiple palindromes, expecting the first
		{"cbbd", "bb"},                 // even-length palindrome
		{"a", "a"},                     // single character
		{"", ""},                       // empty string
		{"ab", "a"},                    // two characters
		{"racecar", "racecar"},         // entire string is a palindrome
		{"abcde", "a"},                 // no repeating characters, first character expected
		{"abacdfgdcaba", "aba"},        // multiple short palindromes
		{"abb", "bb"},                  // two-character palindrome at the end
		{"xyzzyx", "xyzzyx"},           // even-length palindrome over the entire string
		{"madamimadam", "madamimadam"}, // complex palindrome
	}
	for _, test := range tests {
		longestPalindrome := FindLongestPalindrome(test.input)
		if longestPalindrome != test.expected {
			t.Errorf("FindLongestPalindrome(%s) = %s, expected %s", test.input, longestPalindrome, test.expected)
		}
	}
}

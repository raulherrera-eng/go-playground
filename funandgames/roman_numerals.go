package funandgames

// RomanDigit represents a Roman numeral with its decimal and symbol representation
type RomanDigit struct {
	Decimal int
	Symbol  string
}

// Define Roman numerals as constants of type RomanDigit
var (
	I  = RomanDigit{Decimal: 1, Symbol: "I"}
	IV = RomanDigit{Decimal: 4, Symbol: "IV"}
	V  = RomanDigit{Decimal: 5, Symbol: "V"}
	IX = RomanDigit{Decimal: 9, Symbol: "IX"}
	X  = RomanDigit{Decimal: 10, Symbol: "X"}
	XL = RomanDigit{Decimal: 40, Symbol: "XL"}
	L  = RomanDigit{Decimal: 50, Symbol: "L"}
	XC = RomanDigit{Decimal: 90, Symbol: "XC"}
	C  = RomanDigit{Decimal: 100, Symbol: "C"}
	CD = RomanDigit{Decimal: 400, Symbol: "CD"}
	D  = RomanDigit{Decimal: 500, Symbol: "D"}
	CM = RomanDigit{Decimal: 900, Symbol: "CM"}
	M  = RomanDigit{Decimal: 1000, Symbol: "M"}
)

// Sorted in decreasing order
var sortedRomanDigits = []RomanDigit{M, CM, D, CD, C, XC, L, XL, X, IX, V, IV, I}

func GetRomanNumber(num int) string {
	return getRomanNumber(num, "")
}

// Convert a decimal number into
func getRomanNumber(num int, roman string) string {
	if num <= 0 {
		return roman
	}

	next := getNextSmallestRomanDigit(num)
	roman = roman + next.Symbol
	return getRomanNumber(num-next.Decimal, roman)
}

// Get the largest Roman numeral that fits into the given number
func getNextSmallestRomanDigit(num int) RomanDigit {
	for _, roman := range sortedRomanDigits {
		if roman.Decimal <= num {
			return roman
		}
	}
	return I // should not happen
}

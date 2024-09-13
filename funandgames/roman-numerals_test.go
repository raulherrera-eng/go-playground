package funandgames

import "testing"

func TestGetRomanNumber(t *testing.T) {

	tests := []struct {
		decimal int
		roman   string
	}{
		{1, "I"},
		{3, "III"},
		{4, "IV"},
		{6, "VI"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{19, "XIX"},
		{40, "XL"},
		{49, "XLIX"},
		{50, "L"},
		{58, "LVIII"},
		{90, "XC"},
		{99, "XCIX"},
		{100, "C"},
		{400, "CD"},
		{444, "CDXLIV"},
		{500, "D"},
		{501, "DI"},
		{888, "DCCCLXXXVIII"},
		{900, "CM"},
		{1000, "M"},
		{1776, "MDCCLXXVI"},
		{1987, "MCMLXXXVII"},
		{2023, "MMXXIII"},
		{3999, "MMMCMXCIX"},
	}

	for _, tt := range tests {
		var actual = GetRomanNumber(tt.decimal)
		if actual != tt.roman {
			t.Error("Expected", tt.roman, "got", actual)
		}
	}

}

package funandgames

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		input  int
		expect string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},                // The description of "1211" is "one 1, one 2, two 1s".
		{6, "312211"},                // The description of "111221" is "three 1s, two 2s, one 1".
		{7, "13112221"},              // The description of "312211" is "one 3, one 1, two 2s, two 1s".
		{8, "1113213211"},            // The description of "13112221" is "one 1, one 3, two 1s, two 2s, one 1".
		{9, "31131211131221"},        // Continuing for n=9
		{10, "13211311123113112211"}, // Continuing for n=10
	}

	for _, test := range tests {
		if got := CountAndSay(test.input); got != test.expect {
			t.Errorf("countAndSay(%d) = %s, want %s", test.input, got, test.expect)
		}
	}
}

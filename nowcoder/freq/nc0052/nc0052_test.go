package nc0052

import "testing"

func TestNC0052(t *testing.T) {
	var cases = []struct {
		input  string
		output bool
	}{
		{
			input:  "[",
			output: false,
		},
		{
			input:  "[]",
			output: true,
		},
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "([)]",
			output: false,
		},
	}

	for i, c := range cases {
		expected := isValid(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

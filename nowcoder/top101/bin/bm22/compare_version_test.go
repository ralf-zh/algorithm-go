package bm22

import "testing"

func TestCompareVersion(t *testing.T) {
	var cases = []struct {
		input  [2]string
		output int
	}{
		{
			input:  [...]string{"1.1", "1.01"},
			output: 0,
		},
		{
			input:  [...]string{"1.1", "1.1.1"},
			output: -1,
		},
		{
			input:  [...]string{"2.0.1", "2"},
			output: 1,
		},
		{
			input:  [...]string{"0.226", "0.36"},
			output: 1,
		},
	}

	for i, c := range cases {
		expected := compare(c.input[0], c.input[1])
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

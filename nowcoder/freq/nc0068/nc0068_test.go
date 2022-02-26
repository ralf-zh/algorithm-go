package nc0068

import "testing"

func TestNC0068(t *testing.T) {
	var cases = []struct {
		input  int
		output int
	}{
		{
			input:  2,
			output: 2,
		},
		{
			input:  7,
			output: 21,
		},
	}

	for i, c := range cases {
		expected := jumpFloor(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

package nc0001

import "testing"

func TestNC0001(t *testing.T) {
	var cases = []struct {
		s      string
		t      string
		output string
	}{
		{
			s:      "1",
			t:      "99",
			output: "100",
		},
		{
			s:      "114514",
			t:      "",
			output: "114514",
		},
	}

	for i, c := range cases {
		expected := solve(c.s, c.t)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

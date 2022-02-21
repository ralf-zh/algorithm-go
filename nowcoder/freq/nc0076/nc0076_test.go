package nc0076

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNC0076(t *testing.T) {

	var cases = []struct {
		input  []string
		output []int
	}{
		{
			input:  []string{"PSH1", "PSH2", "POP", "POP"},
			output: []int{1, 2},
		},
		{
			input:  []string{"PSH2", "POP", "PSH1", "POP"},
			output: []int{2, 1},
		},
	}

	for i, c := range cases {
		Init()
		var expected []int
		for _, s := range c.input {
			op := s[:3]
			switch op {
			case "PSH":
				n, _ := strconv.Atoi(s[3:])
				Push(n)
			case "POP":
				expected = append(expected, Pop())
			}
		}
		if !reflect.DeepEqual(expected, c.output) {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

func Init() {
	stack1 = nil
	stack2 = nil
}

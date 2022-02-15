package nc0140

import (
	"reflect"
	"testing"
)

func TestNC0140_MySort(t *testing.T) {

	var testSuites = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{5, 2, 3, 1, 4},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{5, 1, 6, 2, 5},
			output: []int{1, 2, 5, 5, 6},
		},
	}

	for i, suite := range testSuites {
		expeced := MySort(suite.input)
		if !reflect.DeepEqual(expeced, suite.output) {
			t.Errorf("[%d] fails, %v != %v", i, expeced, suite.output)
		}
	}
}

func TestNC0140_MySort_Bubble(t *testing.T) {

	var testSuites = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{5, 2, 3, 1, 4},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{5, 1, 6, 2, 5},
			output: []int{1, 2, 5, 5, 6},
		},
	}

	for i, suite := range testSuites {
		expeced := MySort_Bubble(suite.input)
		if !reflect.DeepEqual(expeced, suite.output) {
			t.Errorf("[%d] fails, %v != %v", i, expeced, suite.output)
		}
	}
}

func TestNC0140_MySort_Quick(t *testing.T) {

	var testSuites = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{5, 2, 3, 1, 4},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{5, 1, 6, 2, 5},
			output: []int{1, 2, 5, 5, 6},
		},
	}

	for i, suite := range testSuites {
		expeced := MySort_Quick(suite.input)
		if !reflect.DeepEqual(expeced, suite.output) {
			t.Errorf("[%d] fails, %v != %v", i, expeced, suite.output)
		}
	}
}

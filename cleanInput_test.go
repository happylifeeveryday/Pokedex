package main

import(
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello","world"},
		},
		{
			input: "How are you",
			expected: []string{"how","are","you"},
		},
	}
	for _,c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch：input '%s' expected length %d actual length %d", 
				c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s', expected '%v' but got '%v'",c.input,c.expected,actual)
			}
		}
	}
}
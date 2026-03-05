package internal

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "hello world",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Unique",
			input:    "power",
			expected: []string{"power"},
		},
		{
			name:     "Capitalize",
			input:    "London Paris Berlin",
			expected: []string{"london", "paris", "berlin"},
		},
		{
			name:     "List Upper and Capitalize",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%s: expected %v, got %v", c.name, c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s: expected %v, got %v", c.name, c.expected, actual)
				break
			}
		}
	}
}

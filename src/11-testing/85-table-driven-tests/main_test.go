package main

import "testing"

func TestRemoveNewLineSuffix_Empty(t *testing.T) {
	got := removeNewLineSuffixes("")
	expected := ""
	if got != expected {
		t.Errorf("got: %s", got)
	}
}

func TestRemoveNewLineSuffix_EndingWithCarriageReturnNewLine(t *testing.T) {
	got := removeNewLineSuffixes("a\r\n")
	expected := "a"
	if got != expected {
		t.Errorf("got: %s", got)
	}
}

func TestRemoveNewLineSuffix_EndingWithNewLine(t *testing.T) {
	got := removeNewLineSuffixes("a\n")
	expected := "a"
	if got != expected {
		t.Errorf("got: %s", got)
	}
}

func TestRemoveNewLineSuffix_EndingWithMultipleNewLines(t *testing.T) {
	got := removeNewLineSuffixes("a\n\n\n")
	expected := "a"
	if got != expected {
		t.Errorf("got: %s", got)
	}
}

func TestRemoveNewLineSuffix_EndingWithoutNewLine(t *testing.T) {
	got := removeNewLineSuffixes("a\n")
	expected := "a"
	if got != expected {
		t.Errorf("got: %s", got)
	}
}

func TestRemoveNewLineSuffix(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		`empty`: {
			input:    "",
			expected: "",
		},
		`ending with \r\n`: {
			input:    "a\r\n",
			expected: "a",
		},
		`ending with \n`: {
			input:    "a\n",
			expected: "a",
		},
		`ending with multiple \n`: {
			input:    "a\n\n\n",
			expected: "a",
		},
		`ending without newline`: {
			input:    "a",
			expected: "a",
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := removeNewLineSuffixes(tt.input)
			if got != tt.expected {
				t.Errorf("got: %s, expected: %s", got, tt.expected)
			}
		})
	}
}

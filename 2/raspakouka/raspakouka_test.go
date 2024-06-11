package raspakouka

import "testing"

func TestRaspokouka(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		// {"abcd", "abcd", false},
		// {"45", "", true},
		// {"", "", false},
		// {"qwe\\4\\5", "qwe45", false},
		// {"qwe\\45", "qwe44444", false},
		// {"qwe\\\\5", "qwe\\\\\\", false},
		// {"a\\", "", true},
		// {"a3\\4b", "aaab", false},
	}

	// ● "a4bc2d5e" => "aaaabccddddde"
	// ● "abcd" => "abcd"
	// ● "45" => "" (некорректная строка)
	// ● "" => ""

	for _, test := range tests {
		result, err := Raspokouka(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Unpack(%q) error = %v, wantErr %v", test.input, err, test.hasError)
			continue
		}
		if result != test.expected {
			t.Errorf("Unpack(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}

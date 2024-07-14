package applicant_test

import (
	"movie-matcher/internal/applicant"
	"strings"
	"testing"
)

func TestParseApplicantNameErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        string
		errorMessage string
	}{
		{
			name:         "Empty name",
			input:        "",
			errorMessage: "name cannot be empty. got: ''",
		},
		{
			name:         "1 rune name",
			input:        "a",
			errorMessage: "name must be between 2 and 256 characters. got string of length: 1",
		},
		{
			name:         "Name too long",
			input:        strings.Repeat("a", 257),
			errorMessage: "name must be between 2 and 256 characters. got string of length: 257",
		},
		{
			name:         "Name with only forbidden character",
			input:        "/()'\"<>\\{}",
			errorMessage: "name contains forbidden characters. got: '/()'\"<>\\{}'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := applicant.ParseApplicantName(test.input)
			if err == nil {
				t.Fatalf("expected error to be %s, got nil", test.errorMessage)
			}
			if err.Error() != test.errorMessage {
				t.Fatalf("expected error to be %s, got %s", test.errorMessage, err.Error())
			}
		})
	}
}

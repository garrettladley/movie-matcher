package applicant

import "testing"

func TestNUEmailWithNumberWorks(t *testing.T) {
	t.Parallel()

	email := "ladley.g1@northeastern.edu"
	nuEmail, err := ParseNUEmail(email)
	if err != nil {
		t.Fatalf("expected no error, got %s", err.Error())
	}

	if nuEmail.String() != email {
		t.Fatalf("expected %s, got %s", email, nuEmail.String())
	}
}

func TestNUEmailWithoutNumberWorks(t *testing.T) {
	t.Parallel()

	email := "ladley.g@northeastern.edu"
	nuEmail, err := ParseNUEmail(email)
	if err != nil {
		t.Fatalf("expected no error, got %s", err.Error())
	}

	if nuEmail.String() != email {
		t.Fatalf("expected %s, got %s", email, nuEmail.String())
	}
}

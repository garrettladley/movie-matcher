package types

import "movie-matcher/internal/applicant"

type NotFound struct {
	email applicant.NUEmail
}

package types

import "movie-matcher/internal/applicant"

type NotFound struct {
	NUID applicant.NUID
}

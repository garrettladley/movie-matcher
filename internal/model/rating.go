package model

type Rating string

const (
	RatingG        Rating = "G"
	RatingPG       Rating = "PG"
	RatingPG13     Rating = "PG-13"
	RatingR        Rating = "R"
	RatingNC17     Rating = "NC-17"
	RatingUnrated  Rating = "Unrated"
	RatingApproved Rating = "Approved"
	RatingNA       Rating = "N/A"
)

package model

import "time"

type Person struct {
	Name           string                    `json:"name"`
	Runtime        Preference[time.Duration] `json:"runtime,omitempty"`
	Language       Preference[string]        `json:"language,omitempty"`
	Actor          Preference[string]        `json:"actor,omitempty"`
	RottenTomatoes Preference[int]           `json:"rotten_tomatoes,omitempty"`
	Year           Preference[int]           `json:"year,omitempty"`
	Rating         Preference[string]        `json:"rating,omitempty"`
}

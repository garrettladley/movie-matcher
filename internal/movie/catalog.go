package movie

import (
	"time"
)

var Catalog = []Movie{
	// whiplash
	NewMovieBuilder().
		ID("tt2582802").
		Runtime(time.Duration(106) * time.Minute).
		Languages([]Language{LanguageEnglish}).
		Actors([]string{"Miles Teller", "J.K. Simmons", "Melissa Benoist"}).
		RottenTomatoes(94).
		Year(2014).
		Rating(RatingR).
		Build(),

	// dune 2
	NewMovieBuilder().
		ID("tt15239678").
		Runtime(time.Duration(166) * time.Minute).
		Languages([]Language{LanguageEnglish}).
		Actors([]string{"Timothée Chalamet", "Zendaya", "Rebecca Ferguson"}).
		RottenTomatoes(92).
		Year(2024).
		Rating(RatingPG13).
		Build(),

	// lalaland
	NewMovieBuilder().
		ID("tt3783958").
		Runtime(time.Duration(128) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageCantonese}).
		Actors([]string{"Ryan Gosling", "Emma Stone", "Rosemarie DeWitt"}).
		RottenTomatoes(91).
		Year(2016).
		Rating(RatingPG13).
		Build(),

	// the lego movie
	NewMovieBuilder().
		ID("tt1490017").
		Runtime(time.Duration(100) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageTurkish}).
		Actors([]string{"Chris Pratt", "Will Ferrell", "Elizabeth Banks"}).
		RottenTomatoes(96).
		Year(2014).
		Rating(RatingPG).
		Build(),

	// 2001: a space odyssey
	NewMovieBuilder().
		ID("tt0062622").
		Runtime(time.Duration(149) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageRussian, LanguageFrench}).
		Actors([]string{"Keir Dullea", "Gary Lockwood", "William Sylvester"}).
		RottenTomatoes(92).
		Year(1968).
		Rating(RatingG).
		Build(),

	// inside out 2
	NewMovieBuilder().
		ID("tt22022452").
		Runtime(time.Duration(95) * time.Minute).
		Languages([]Language{LanguageEnglish}).
		Actors([]string{"Amy Poehler", "Maya Hawke", "Kensington Tallman"}).
		RottenTomatoes(92).
		Year(2024).
		Rating(RatingPG).
		Build(),

	// the grand budapest hotel
	NewMovieBuilder().
		ID("tt2278388").
		Runtime(time.Duration(99) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageFrench, LanguageGerman}).
		Actors([]string{"Ralph Fiennes", "F. Murray Abraham", "Mathieu Amalric"}).
		RottenTomatoes(92).
		Year(2014).
		Rating(RatingR).
		Build(),

	// the imitation game
	NewMovieBuilder().
		ID("tt2084970").
		Runtime(time.Duration(114) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageGerman}).
		Actors([]string{"Benedict Cumberbatch", "Keira Knightley", "Matthew Goode"}).
		RottenTomatoes(90).
		Year(2014).
		Rating(RatingPG13).
		Build(),

	// apollo 13
	NewMovieBuilder().
		ID("tt0112384").
		Runtime(time.Duration(140) * time.Minute).
		Languages([]Language{LanguageEnglish}).
		Actors([]string{"Tom Hanks", "Bill Paxton", "Kevin Bacon"}).
		RottenTomatoes(96).
		Year(1995).
		Rating(RatingPG).
		Build(),

	// catch me if you can
	NewMovieBuilder().
		ID("tt0264464").
		Runtime(time.Duration(141) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageFrench}).
		Actors([]string{"Leonardo DiCaprio", "Tom Hanks", "Christopher Walken"}).
		RottenTomatoes(96).
		Year(2002).
		Rating(RatingPG13).
		Build(),

	// the fantastic mr. fox
	NewMovieBuilder().
		ID("tt0432283").
		Runtime(time.Duration(87) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageFrench}).
		Actors([]string{"George Clooney", "Meryl Streep", "Bill Murray"}).
		RottenTomatoes(93).
		Year(2009).
		Rating(RatingPG).
		Build(),

	// the minions movie
	NewMovieBuilder().
		ID("tt2293640").
		Runtime(time.Duration(91) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageSpanish, LanguageItalian}).
		Actors([]string{"Sandra Bullock", "Jon Hamm", "Michael Keaton"}).
		RottenTomatoes(56).
		Year(2015).
		Rating(RatingPG).
		Build(),

	// goldfinger
	NewMovieBuilder().
		ID("tt0058150").
		Runtime(time.Duration(110) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageChinese, LanguageSpanish}).
		Actors([]string{"Sean Connery", "Gert Frobe", "Honor Blackman"}).
		RottenTomatoes(99).
		Year(1964).
		Build(),

	// skyfall
	NewMovieBuilder().
		ID("tt1074638").
		Runtime(time.Duration(143) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageTurkish, LanguageShanghainese, LanguagePortuguese, LanguageJapanese}).
		Actors([]string{"Daniel Craig", "Javier Bardem", "Naomie Harris"}).
		RottenTomatoes(92).
		Year(2012).
		Rating(RatingPG13).
		Build(),

	// the social network
	NewMovieBuilder().
		ID("tt1285016").
		Runtime(time.Duration(120) * time.Minute).
		Languages([]Language{LanguageEnglish, LanguageFrench}).
		Actors([]string{"Jesse Eisenberg", "Andrew Garfield", "Justin Timberlake"}).
		RottenTomatoes(96).
		Year(2010).
		Rating(RatingPG13).
		Build(),
}

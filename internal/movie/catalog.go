package movie

import "movie-matcher/internal/ordered_set"

var Catalog = ordered_set.New[ID](
	"tt2582802",  // whiplash
	"tt15239678", // dune 2
	"tt3783958",  // lalaland
	"tt1490017",  // the lego movie
	"tt0062622",  // 2001: a space odyssey
	"tt22022452", // inside out 2
	"tt2278388",  // the grand budapest hotel
	"tt2084970",  // the imitation game
	"tt0112384",  // apollo 13
	"tt0264464",  // catch me if you can
	"tt0432283",  // fantastic mr fox
	"tt2293640",  // the minions movie
	"tt0058150",  // goldfinger
	"tt1074638",  // skyfall
	"tt1285016",  // the social network
)

var TopMoviesCatalog = ordered_set.New[ID](
	"tt2582802",  // whiplash
	"tt15239678", // dune 2
	"tt3783958",  // lalaland
	"tt1490017",  // the lego movie
	"tt0062622",  // 2001: a space odyssey

)

var RecommendationsCatalog = ordered_set.New[ID](
	"tt22022452", // inside out 2
	"tt2278388",  // the grand budapest hotel
	"tt2084970",  // the imitation game
	"tt0112384",  // apollo 13
	"tt0264464",  // catch me if you can
)

var ContinueWatchingCatalog = ordered_set.New[ID](
	"tt0432283", // fantastic mr fox
	"tt2293640", // the minions movie
	"tt0058150", // goldfinger
	"tt1074638", // skyfall
	"tt1285016", // the social network
)

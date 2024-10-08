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

var PosterCatalog = map[ID]string{
	"tt2582802":  "https://i.imgur.com/PlNisIp.jpg", // whiplash
	"tt15239678": "https://i.imgur.com/UQHlMQ1.jpg", // dune 2
	"tt3783958":  "https://i.imgur.com/8YmryZL.jpg", // lalaland
	"tt1490017":  "https://i.imgur.com/nuyD2fN.jpg", // the lego movie
	"tt0062622":  "https://i.imgur.com/cqEE5PG.jpg", // 2001: a space odyssey
	"tt22022452": "https://i.imgur.com/nvzrkkz.jpg", // inside out 2
	"tt2278388":  "https://i.imgur.com/1mM13lC.jpg", // the grand budapest hotel
	"tt2084970":  "https://i.imgur.com/cKz9dhr.jpg", // the imitation game
	"tt0112384":  "https://i.imgur.com/DIdS6mp.jpg", // apollo 13
	"tt0264464":  "https://i.imgur.com/qGqain2.jpg", // catch me if you can
	"tt0432283":  "https://i.imgur.com/PDpAJLK.jpg", // fantastic mr fox
	"tt2293640":  "https://i.imgur.com/pPkUfRP.jpg", // the minions movie
	"tt0058150":  "https://i.imgur.com/VWKIDR0.jpg", // goldfinger
	"tt1074638":  "https://i.imgur.com/XHeEjqy.jpg", // skyfall
	"tt1285016":  "https://i.imgur.com/q0WxmVj.jpg", // the social network
}

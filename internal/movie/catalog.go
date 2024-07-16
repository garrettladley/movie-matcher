package movie

var Catalog = []ID{
	"tt1216475",  // whiplash
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
}

// FRONTEND CHALLANGE CATALOGS:
var CatalogDisplay = []struct {
    ID                 string
    MovieDisplayDetails MovieDisplayDetails
}{
    {"tt1216475", MovieDisplayDetails{"https://i.imgur.com/PlNisIp.jpeg", nil}}, // whiplash
    {"tt15239678", MovieDisplayDetails{"https://i.imgur.com/UQHlMQ1.jpeg", nil}}, // dune 2
    {"tt3783958", MovieDisplayDetails{"https://i.imgur.com/8YmryZL.jpeg", nil}}, // lalaland
    {"tt1490017", MovieDisplayDetails{"https://i.imgur.com/nuyD2fN.jpeg", nil}}, // the lego movie
    {"tt0062622", MovieDisplayDetails{"https://i.imgur.com/cqEE5PG.jpeg", nil}}, // 2001: a space odyssey
    {"tt22022452", MovieDisplayDetails{"https://i.imgur.com/nvzrkkz.jpeg", nil}}, // inside out 2
    {"tt2278388", MovieDisplayDetails{"https://i.imgur.com/1mM13lC.jpeg", nil}}, // the grand budapest hotel
    {"tt2084970", MovieDisplayDetails{"https://i.imgur.com/cKz9dhr.jpeg", nil}}, // the imitation game
    {"tt0112384", MovieDisplayDetails{"https://i.imgur.com/DIdS6mp.jpeg", nil}}, // apollo 13
    {"tt0264464", MovieDisplayDetails{"https://i.imgur.com/qGqain2.jpeg", nil}}, // catch me if you can
    {"tt0432283", MovieDisplayDetails{"https://i.imgur.com/PDpAJLK.jpeg", nil}}, // fantastic mr fox
    {"tt2293640", MovieDisplayDetails{"https://i.imgur.com/pPkUfRP.jpeg", nil}}, // the minions movie
    {"tt0058150", MovieDisplayDetails{"https://i.imgur.com/VWKIDR0.jpeg", nil}}, // goldfinger
    {"tt1074638", MovieDisplayDetails{"https://i.imgur.com/XHeEjqy.jpeg", nil}}, // skyfall
    {"tt1285016", MovieDisplayDetails{"https://i.imgur.com/q0WxmVj.jpeg", nil}}, // the social network
}

package pref_gen

import "movie-matcher/internal/movie"

var names = [...]string{
	"Harry", "Luke", "Frodo", "Sherlock", "Tony",
	"Jon", "Walter", "Arya", "Forrest", "James",
	"Indiana", "Buffy", "Bruce", "Ellen", "Jack",
	"Michael", "Homer", "Marty", "Katniss", "Rocky",
	"Han", "Rick", "Daenerys", "Sarah", "Gandalf",
	"Peter", "Hermione", "Bilbo", "John", "Natasha",
}

var ratings = [...]movie.Rating{
	movie.RatingG, movie.RatingPG, movie.RatingPG13, movie.RatingR, movie.RatingNC17,
}

var genres = [...]string{
	"Action", "Adventure", "Animation", "Biography", "Comedy",
	"Crime", "Documentary", "Drama", "Family", "Fantasy",
	"Film-Noir", "History", "Horror", "Music", "Musical",
	"Mystery", "Romance", "Sci-Fi", "Sport", "Thriller",
	"War", "Western",
}

var directors = [...]string{
	"Steven Spielberg", "Martin Scorsese", "Christopher Nolan", "Quentin Tarantino", "Alfred Hitchcock",
	"Stanley Kubrick", "James Cameron", "Francis Ford Coppola", "Ridley Scott", "Peter Jackson",
	"David Fincher", "Tim Burton", "Clint Eastwood", "Coen Brothers", "Woody Allen",
}

var actors = [...]string{
	"Marlon Brando", "Robert De Niro", "Al Pacino", "Leonardo DiCaprio", "Tom Hanks",
	"Johnny Depp", "Denzel Washington", "Jack Nicholson", "Morgan Freeman", "Brad Pitt",
	"Anthony Hopkins", "Heath Ledger", "Daniel Day-Lewis", "Gary Oldman", "Robert Downey Jr.",
}

var plotElements = [...]string{
	"love", "war", "friendship", "betrayal", "revenge",
	"mystery", "discovery", "journey", "tragedy", "escape",
	"sacrifice", "romance", "quest", "survival", "forgiveness",
}

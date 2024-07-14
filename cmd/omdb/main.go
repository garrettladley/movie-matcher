package main

import (
	"context"
	"fmt"
	"movie-matcher/internal/services/omdb"

	go_json "github.com/goccy/go-json"

	"github.com/alecthomas/kong"
)

var cli struct {
	Find FindCmd `cmd:"" help:"Find a specific movie by its title or IMDb movie ID"`
}

type FindCmd struct {
	Title string `xor:"xorGroup" required:"" help:"The movie's title"`
	ID    string `xor:"xorGroup" required:"" help:"The movie's IMDb ID"`
}

func (cmd *FindCmd) Run() error {
	var (
		movie omdb.Movie
		err   error
	)

	if cmd.Title != "" {
		movie, err = omdb.FindMovieByTitle(context.Background(), cmd.Title)
	} else if cmd.ID != "" {
		movie, err = omdb.FindMovieById(context.Background(), cmd.ID)
	}

	if err != nil {
		return err
	}

	data, err := go_json.Marshal(movie)
	if err != nil {
		return nil
	}
	_, err = fmt.Printf("%s\n", string(data))
	return err
}

func main() {
	ctx := kong.Parse(&cli)
	ctx.FatalIfErrorf(ctx.Run())
}

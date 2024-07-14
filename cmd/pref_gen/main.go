package main

import (
	"fmt"
	"math/rand"
	"movie-matcher/internal/services/pref_gen"
	"time"

	go_json "github.com/goccy/go-json"

	"github.com/alecthomas/kong"
)

var cli struct {
	Generate GenerateCmd `cmd:"" help:"Generate the specified number of people and their preferences"`
}

type GenerateCmd struct {
	Count uint `arg:"" required:"" help:"The number of people to generate"`
}

func (cmd *GenerateCmd) Run() error {
	people := pref_gen.GeneratePeople(rand.New(rand.NewSource(time.Now().UnixNano())), cmd.Count)
	for _, person := range people {
		data, err := go_json.Marshal(person)
		if err != nil {
			return err
		}
		_, err = fmt.Printf("%s\n", string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ctx := kong.Parse(&cli)
	ctx.FatalIfErrorf(ctx.Run())
}

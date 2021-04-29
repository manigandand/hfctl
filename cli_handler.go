package main

import (
	"encoding/json"
	"hfctl/types"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// defaultStatsCalHandler command handler calculates the recipe stats from the given fixtures
// for the default parameters
// Assumption: we expect the user always provide filepath or urlpath to the fixtures
// else, load default test fixtures from the fs
//
// TODO: the given default test fixtures is more than 100MB,
// so rethink to adding that into the docker image
func defaultStatsCalHandler(c *cli.Context) error {
	fileReader, jsonDecoder, err := readRecipeData(c.String("file"))
	defer func() { fileReader.Close() }()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return printStats(NewDefaultStatsInput(jsonDecoder).Calculate())
}

// recipeSearchStats command handler calculates the recipe stats from the given fixtures
// for the given search parameters like recipe_names, postcode & time_window
func recipeSearchStats(c *cli.Context) error {
	postcode := c.String("postcode")
	if ts(postcode) == "" {
		postcode = defaultPostcode
	}
	timeWindow := c.String("time-window")
	if ts(timeWindow) == "" {
		timeWindow = defaultTimeWindow
	}
	recipes := c.String("recipes")
	if ts(recipes) == "" {
		recipes = defaultRecipeToSearch
	}

	fileReader, jsonDecoder, err := readRecipeData(c.String("file"))
	defer func() { fileReader.Close() }()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	stats := NewStats(jsonDecoder, postcode, timeWindow, recipes).Calculate()

	return printStats(stats)
}

func printStats(recipeStats *types.RecipeStats) error {
	statsByte, err := json.MarshalIndent(recipeStats, "", " ")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	olog.Println(string(statsByte))
	return nil
}

// readRecipeData decodes the recipe fixtures data from the given filePath or URL
func readRecipeData(f string) (*os.File, *json.Decoder, error) {
	if ts(f) == "" {
		// load default test data, if the filepath not given
		f = "test/json_file.log"
	}
	fileReader, err := os.Open(ts(f))
	if err != nil {
		return nil, nil, err
	}
	return fileReader, json.NewDecoder(fileReader), nil
}

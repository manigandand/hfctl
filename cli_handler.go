package main

import (
	"encoding/json"
	"errors"
	"hfctl/types"
	"hfctl/utils"
	"log"

	"github.com/urfave/cli/v2"
)

// defaultStatsCalHandler command handler calculates the recipe stats from the given fixtures
// for the default parameters
// Assumption: we expect the user always provide filepath or urlpath to the fixtures
//
// TODO: load default fixtures from the fs, the given default test fixtures
// is more than 100MB, so rethink to adding that into the docker image
func defaultStatsCalHandler(c *cli.Context) error {
	data, err := readRecipeData(c.String("file"))
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return printStats(NewDefaultStatsInput(data).Calculate())
}

// recipeSearchStats command handler calculates the recipe stats from the given fixtures
// for the given search parameters like recipe_names, postcode & time_window
func recipeSearchStats(c *cli.Context) error {
	return nil
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
func readRecipeData(f string) ([]*types.Recipe, error) {
	var data []*types.Recipe
	if ts(f) == "" {
		return nil, errors.New("please specify file path")
	}
	body, err := utils.GetBytesForFileOrURL(ts(f))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

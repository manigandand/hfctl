package main

import (
	"encoding/json"
	"hfctl/types"
	"hfctl/utils"
	"io"
	"regexp"
	"sort"
	"strings"
)

var ts = strings.TrimSpace

const (
	defaultPostcode       string = "10120"
	defaultTimeWindow     string = "10AM - 2PM"
	defaultRecipeToSearch string = "Potato,Veggie,Mushroom"
)

// recipe stats calculator inputs
type recipeStatsInput struct {
	data       *json.Decoder
	postcode   string
	timeWindow string
	recipes    []string
}

// NewDefaultStatsInput returns the recipe stats calculator inputs object with
// the default inputs
func NewDefaultStatsInput(jsonDecoder *json.Decoder) *recipeStatsInput {
	return &recipeStatsInput{
		data:       jsonDecoder,
		postcode:   defaultPostcode,
		timeWindow: defaultTimeWindow,
		recipes:    strings.Split(defaultRecipeToSearch, ","),
	}
}

// NewStats returns the recipe stats calculator inputs object with the given inputs
func NewStats(jsonDecoder *json.Decoder, postcode, timeWindow, recipes string) *recipeStatsInput {
	return &recipeStatsInput{
		data:       jsonDecoder,
		postcode:   postcode,
		timeWindow: timeWindow,
		recipes:    strings.Split(recipes, ","),
	}
}

// Calculate - calculates the recipe stats for the given scinarios
// 1. Count the number of unique recipe names.
// 2. Count the number of occurences for each unique recipe name (alphabetically ordered by recipe name).
// 3. Find the postcode with most delivered recipes.
// 4. Count the number of deliveries to postcode `10120` that lie within the delivery time between 10AM and 3PM.
// 5. List the recipe names (alphabetically ordered) that contain in their name one of the following words:
// Potato, Veggie, Mushroom
//
// returns the final object
func (i *recipeStatsInput) Calculate() *types.RecipeStats {
	var (
		postcodeDeliveryCount int
	)
	matchedRecipes := make([]string, 0)
	uniqueRecipeCount := make(types.RecipeCount)
	deliveryCountByPostcode := make(types.PostcodeDeliveryCount)
	matchedRecipe := make(map[string]bool)

	inputTime := parseInputTime(i.timeWindow)

	for {
		// decode data from the json decoder
		var d *types.Recipe
		if err := i.data.Decode(&d); err == io.EOF {
			break
		}
		if d == nil {
			continue
		}
		d.Recipe = ts(d.Recipe)
		// unique recipe count
		uniqueRecipeCount.Inc(d.Recipe)

		// delivery count by postalcode
		deliveryCountByPostcode.Inc(d.Postcode)

		// delivery count for postalcode
		dt := parseDeliveryTime(d.Delivery)
		if dt != nil && inputTime != nil {
			if d.Postcode == i.postcode && (inputTime.From >= dt.From && inputTime.To <= dt.To) {
				postcodeDeliveryCount++
			}
		}

		// match recipe with any one of the given recipe inputs
		for _, r := range i.recipes {
			if matched, _ := regexp.MatchString(ts(r), d.Recipe); matched {
				matchedRecipe[d.Recipe] = true
				break
			}
		}
	}
	for r := range matchedRecipe {
		matchedRecipes = append(matchedRecipes, r)
	}
	sort.Strings(matchedRecipes)

	// final recipe stats cal
	stats := &types.RecipeStats{
		UniqueRecipeCount: len(uniqueRecipeCount),
		CountPerRecipe:    uniqueRecipeCount.ToStruct(),
		BusiestPostcode:   deliveryCountByPostcode.GetBusiestPostcode(),
		DeliveryStats: types.DeliveryStatsByPostcode{
			Postcode:      i.postcode,
			From:          inputTime.FromStr,
			To:            inputTime.ToStr,
			DeliveryCount: postcodeDeliveryCount,
		},
		MatchByName: matchedRecipes,
	}

	return stats
}

// parseDeliveryTime parse the deilvery time
// Format: "Tuesday 4AM - 2PM"
func parseDeliveryTime(ts string) *types.TS {
	tl := strings.Split(ts, " ")
	if len(tl) < 1 {
		return nil
	}

	return &types.TS{
		FromStr: tl[1],
		From:    utils.TimeStrToRailwayTime(tl[1]),
		ToStr:   tl[len(tl)-1],
		To:      utils.TimeStrToRailwayTime(tl[len(tl)-1]),
	}
}

// parseInputTime parse the input time
// Format: "10AM - 2PM"
func parseInputTime(ts string) *types.TS {
	tl := strings.Split(ts, " ")
	if len(tl) < 1 {
		return nil
	}

	return &types.TS{
		FromStr: tl[0],
		From:    utils.TimeStrToRailwayTime(tl[0]),
		ToStr:   tl[len(tl)-1],
		To:      utils.TimeStrToRailwayTime(tl[len(tl)-1]),
	}
}

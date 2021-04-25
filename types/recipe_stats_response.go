package types

// RecipeStats holds the final response json type
type RecipeStats struct {
	UniqueRecipeCount int                     `json:"unique_recipe_count"`
	CountPerRecipe    []CountPerRecipe        `json:"count_per_recipe"`
	BusiestPostcode   BusiestPostcode         `json:"busiest_postcode"`
	DeliveryStats     DeliveryStatsByPostcode `json:"count_per_postcode_and_time"`
	MatchByName       []string                `json:"match_by_name"`
}

// CountPerRecipe holds `count_per_recipe` response type
type CountPerRecipe struct {
	Recipe string `json:"recipe"`
	Count  int    `json:"count"`
}

// BusiestPostcode holds `busiest_postcode` response type
type BusiestPostcode struct {
	Postcode      string `json:"postcode"`
	DeliveryCount int    `json:"delivery_count"`
}

// DeliveryStatsByPostcode holds `count_per_postcode_and_time` response type
type DeliveryStatsByPostcode struct {
	Postcode      string `json:"postcode"`
	From          string `json:"from"`
	To            string `json:"to"`
	DeliveryCount int    `json:"delivery_count"`
}

type TS struct {
	FromStr string
	ToStr   string

	From int
	To   int
}

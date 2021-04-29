package types

import "sort"

const (
	MaxDistinctRecipes = 2000
)

// RecipeCount: custom map type which holdes the counts(occurences) of the recipes
type RecipeCount map[string]int

// increment the delivery counts by 1 for the given postal codes
func (rc RecipeCount) Inc(recipe string) {
	if len(recipe) > 100 || len(rc) == MaxDistinctRecipes {
		return
	}
	count, ok := rc[recipe]
	if !ok {
		rc[recipe] = 1
		return
	}
	rc[recipe] = count + 1
}

// converts the map to `CountPerRecipe`
func (rc RecipeCount) ToStruct() []CountPerRecipe {
	uniqueRecipe := make([]CountPerRecipe, 0, len(rc))
	for r, c := range rc {
		uniqueRecipe = append(uniqueRecipe, CountPerRecipe{
			Recipe: r,
			Count:  c,
		})
	}
	sort.Sort(SortByRecipeName(uniqueRecipe))

	return uniqueRecipe
}

// custom sort helper type
type SortByRecipeName []CountPerRecipe

func (a SortByRecipeName) Len() int           { return len(a) }
func (a SortByRecipeName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByRecipeName) Less(i, j int) bool { return a[i].Recipe < a[j].Recipe }

// PostcodeDeliveryCount: custom map type which holdes the delivery counts
// of the postal codes
type PostcodeDeliveryCount map[string]int

// increment the delivery counts by 1 for the given postal codes
func (pdc PostcodeDeliveryCount) Inc(postcode string) {
	if len(postcode) > 10 {
		return
	}
	dCount, ok := pdc[postcode]
	if !ok {
		pdc[postcode] = 1
		return
	}
	pdc[postcode] = dCount + 1
}

// converts the map to `BusiestPostcode`
func (pdc PostcodeDeliveryCount) GetBusiestPostcode() BusiestPostcode {
	busiestPostcodes := make([]BusiestPostcode, 0, len(pdc))
	for p, c := range pdc {
		busiestPostcodes = append(busiestPostcodes, BusiestPostcode{
			Postcode:      p,
			DeliveryCount: c,
		})
	}
	sort.Sort(SortByDeliveryCount(busiestPostcodes))

	if len(busiestPostcodes) == 0 {
		return BusiestPostcode{}
	}

	return busiestPostcodes[0]
}

// custom sort helper type
type SortByDeliveryCount []BusiestPostcode

func (a SortByDeliveryCount) Len() int           { return len(a) }
func (a SortByDeliveryCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByDeliveryCount) Less(i, j int) bool { return a[i].DeliveryCount > a[j].DeliveryCount }

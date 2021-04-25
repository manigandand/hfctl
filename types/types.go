package types

import "sort"

// RecipeCountMap: custom map type which holdes the counts(occurences) of the recipes
type RecipeCountMap map[string]int

// increment the delivery counts by 1 for the given postal codes
func (rcMap RecipeCountMap) Inc(recipe string) {
	count, ok := rcMap[recipe]
	if !ok {
		rcMap[recipe] = 1
		return
	}
	rcMap[recipe] = count + 1
}

// converts the map to `CountPerRecipe`
func (rcMap RecipeCountMap) ToStruct() []CountPerRecipe {
	uniqueRecipe := make([]CountPerRecipe, 0, len(rcMap))
	for r, c := range rcMap {
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

// PostcodeDeliveryCountMap: custom map type which holdes the delivery counts
// of the postal codes
type PostcodeDeliveryCountMap map[string]int

// increment the delivery counts by 1 for the given postal codes
func (pdcMap PostcodeDeliveryCountMap) Inc(postcode string) {
	dCount, ok := pdcMap[postcode]
	if !ok {
		pdcMap[postcode] = 1
		return
	}
	pdcMap[postcode] = dCount + 1
}

// converts the map to `BusiestPostcode`
func (pdcMap PostcodeDeliveryCountMap) GetBusiestPostcode() BusiestPostcode {
	busiestPostcodes := make([]BusiestPostcode, 0, len(pdcMap))
	for p, c := range pdcMap {
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

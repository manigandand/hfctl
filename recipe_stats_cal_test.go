package main

import (
	"hfctl/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	Context("Recipe Stats Calculator test", func() {
		When("not valid data sent", func() {
			var (
				res *types.RecipeStats
			)
			BeforeEach(func() {
			})
			JustBeforeEach(func() {
				res = NewDefaultStatsInput(nil).Calculate()
			})
			It("should stdout the stats", func() {
				Expect(res.UniqueRecipeCount).Should(Equal(0))
				Expect(len(res.CountPerRecipe)).Should(Equal(0))
				Expect(res.BusiestPostcode.Postcode).Should(Equal(""))
				Expect(res.BusiestPostcode.DeliveryCount).Should(Equal(0))
				Expect(res.DeliveryStats.Postcode).Should(Equal("10120"))
				Expect(res.DeliveryStats.DeliveryCount).Should(Equal(0))
				Expect(res.DeliveryStats.From).Should(Equal("10AM"))
				Expect(res.DeliveryStats.To).Should(Equal("2PM"))
				Expect(len(res.MatchByName)).Should(Equal(0))
			})
		})

		When("valid file", func() {
			var (
				data []*types.Recipe
				err  error
				res  *types.RecipeStats
			)
			BeforeEach(func() {
				data, err = readRecipeData("./test/hf_test_calculation_fixtures.json")
				Expect(err).ShouldNot(HaveOccurred())
			})
			JustBeforeEach(func() {
				res = NewDefaultStatsInput(data).Calculate()
			})
			It("should stdout the stats", func() {
				Expect(res.UniqueRecipeCount).Should(Equal(29))
				Expect(len(res.CountPerRecipe)).Should(Equal(29))
				Expect(res.CountPerRecipe[0].Recipe).Should(Equal("Cajun-Spiced Pulled Pork"))
				Expect(res.CountPerRecipe[0].Count).Should(Equal(39))
				Expect(res.BusiestPostcode.Postcode).Should(Equal("10186"))
				Expect(res.BusiestPostcode.DeliveryCount).Should(Equal(11))
				Expect(res.DeliveryStats.Postcode).Should(Equal("10120"))
				Expect(res.DeliveryStats.DeliveryCount).Should(Equal(5))
				Expect(res.DeliveryStats.From).Should(Equal("10AM"))
				Expect(res.DeliveryStats.To).Should(Equal("2PM"))
				Expect(len(res.MatchByName)).Should(Equal(2))
			})
		})
	})
})

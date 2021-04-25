package main

import (
	"hfctl/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLI handler test suite", func() {
	Context("Read Recipe Data From File", func() {
		When("empty file provided", func() {
			var (
				data []*types.Recipe
				err  error
			)
			BeforeEach(func() {
				data, err = readRecipeData("")

			})
			JustBeforeEach(func() {
				// perr = printStats(data, defaultPostcode, defaultTimeWindow, defaultRecipeToSearch)
			})
			It("should return please specify file path", func() {
				Expect(err.Error()).Should(Equal("please specify file path"))
				Expect(data).Should(BeNil())
			})
		})

		When("invalid file provided", func() {
			var (
				data []*types.Recipe
				err  error
			)
			BeforeEach(func() {
				data, err = readRecipeData("./test/invalid.json")

			})
			JustBeforeEach(func() {
				// perr = printStats(data, defaultPostcode, defaultTimeWindow, defaultRecipeToSearch)
			})
			It("should return stat ./test/invalid.json: no such file or directory", func() {
				Expect(err.Error()).Should(Equal("stat ./test/invalid.json: no such file or directory"))
				Expect(data).Should(BeNil())
			})
		})

		When("invalid file provided", func() {
			var (
				data []*types.Recipe
				err  error
			)
			BeforeEach(func() {
				data, err = readRecipeData("./test/invalid.json")

			})
			JustBeforeEach(func() {
				// perr = printStats(data, defaultPostcode, defaultTimeWindow, defaultRecipeToSearch)
			})
			It("should return stat ./test/invalid.json: no such file or directory", func() {
				Expect(err.Error()).Should(Equal("stat ./test/invalid.json: no such file or directory"))
				Expect(data).Should(BeNil())
			})
		})

		When("invalid file content to read", func() {
			var (
				data      []*types.Recipe
				err, perr error
			)
			BeforeEach(func() {
				data, err = readRecipeData("./test/invalid_data.json")
			})
			JustBeforeEach(func() {
				perr = printStats(NewDefaultStatsInput(data).Calculate())
			})
			It("should return stat ./test/invalid.json: no such file or directory", func() {
				Expect(err.Error()).Should(Equal("invalid character ']' looking for beginning of object key string"))
				Expect(data).Should(BeNil())
				Expect(perr).ShouldNot(HaveOccurred())
			})
		})

		When("valid file", func() {
			var (
				data      []*types.Recipe
				err, perr error
			)
			BeforeEach(func() {
				data, err = readRecipeData("./test/hf_test_calculation_fixtures.json")
				Expect(err).ShouldNot(HaveOccurred())
			})
			JustBeforeEach(func() {
				perr = printStats(NewDefaultStatsInput(data).Calculate())
			})
			It("should stdout the stats", func() {
				Expect(perr).ShouldNot(HaveOccurred())
			})
		})
	})
})

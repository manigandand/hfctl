package main

import (
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CLI handler test suite", func() {
	Context("Read Recipe Data From File", func() {
		When("invalid file provided", func() {
			var (
				data *json.Decoder
				f    *os.File
				err  error
			)
			BeforeEach(func() {
				f, data, err = readRecipeData("./test/invalid.json")

			})
			AfterEach(func() {
				f.Close()
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
				f    *os.File
				data *json.Decoder
				err  error
			)
			BeforeEach(func() {
				f, data, err = readRecipeData("./test/invalid.json")

			})
			AfterEach(func() {
				f.Close()
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
				f         *os.File
				data      *json.Decoder
				err, perr error
			)
			BeforeEach(func() {
				f, data, err = readRecipeData("./test/invalid_data.json")
			})
			AfterEach(func() {
				f.Close()
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
				f         *os.File
				data      *json.Decoder
				err, perr error
			)
			BeforeEach(func() {
				f, data, err = readRecipeData("./test/hf_test_calculation_fixtures.json")
				Expect(err).ShouldNot(HaveOccurred())
			})
			AfterEach(func() {
				f.Close()
			})
			JustBeforeEach(func() {
				perr = printStats(NewDefaultStatsInput(data).Calculate())
			})
			It("should stdout the stats", func() {
				Expect(perr).ShouldNot(HaveOccurred())
			})
		})

		When("empty file provided", func() {
			var (
				f         *os.File
				data      *json.Decoder
				err, perr error
			)
			BeforeEach(func() {
				f, data, err = readRecipeData("./test/hf_test_calculation_fixtures.json")
				Expect(err).ShouldNot(HaveOccurred())
			})
			AfterEach(func() {
				f.Close()
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

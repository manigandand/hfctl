package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

// FIXME: known issue: memory & run time spikes up when loading large data sets

var (
	olog *log.Logger
)

func init() {
	log.SetOutput(os.Stdout)
	olog = log.New(os.Stdout, "", log.LstdFlags)
	// don't want to print the date time prefix
	olog.SetFlags(0)
}

// init cli app configurations
func main() {
	app := cli.NewApp()
	// set global flag to get file fixtures
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "Loads the recipe fixtures data from `FILE`",
		},
	}
	// set the search cmd and flags
	app.Commands = []*cli.Command{
		{
			Name:     "search",
			Aliases:  []string{"s"},
			Usage:    "Search recipes by name and count the no.of deliveries per postalcode between time window.",
			Category: "Search Commands",
			Action:   recipeSearchStats,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "recipes",
					Aliases:  []string{"r"},
					Usage:    "Search all the recipes which has these names `RECIPE`",
					Value:    "Potato,Veggie,Mushroom",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "postcode",
					Aliases:  []string{"pin"},
					Usage:    "Get the no.of deliveries count by `POSTAL CODE`",
					Value:    "10120",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "time-window",
					Aliases:  []string{"tw"},
					Usage:    "Get the no.of deliveries count by `POSTAL CODE` between the time window",
					Value:    "10AM - 2PM",
					Required: false,
				},
			},
			CustomHelpTemplate: searchHelpString,
		},
	}
	// init
	app.Authors = []*cli.Author{
		{
			Name: "HelloFresh",
		},
		{
			Name:  "Manigandan Dharmalingam",
			Email: "manigandan.jeff@gmail.com",
		},
	}
	app.Description = `
	Command Line Interface(CLI) for HelloFresh Recipe Stats Calculator.
hfctl gives you a meaningful stats of recipes and let you search by postalcode and recipe names.`
	app.Name = "hfctl"
	app.Usage = "CLI app to calculate HelloFresh Recipe Stats"
	app.UsageText = `hfctl [global options] command [command options] [arguments...]
	eg:
	hfctl -f ./hf_fixtures.json search --recipes="Mango,Chicken" --postcode="10120" --time-window="10AM-2PM"
	`
	app.EnableBashCompletion = true
	app.Version = "v1.0.0"
	app.Action = defaultStatsCalHandler
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

const searchHelpString = `
Search recipes by name and count the no.of deliveries per postalcode between time window.

Examples:
  # Search all the recipes which has these RECIPE names.
  hfctl -f ./test/hf_test_calculation_fixtures.json search --recipes="Potato,Veggie,Mushroom"
  hfctl -f ./test/hf_test_calculation_fixtures.json search -r="Mango,Chicken"

  # Get the no.of deliveries count by POSTAL CODE and TIME WINDOW.
  hfctl -f ./test/hf_test_calculation_fixtures.json search --postcode="10120" --time-window="10AM - 2PM"
  hfctl -f ./test/hf_test_calculation_fixtures.json search -pc="10120" -tw="10AM - 2PM"

Options:
    -r, --recipes='': Search all the recipes which has these RECIPE KEY names.
  	-pin, --postcode='': Get the no.of deliveries count by POSTAL CODE.
  	-tw, --time-window='': Get the no.of deliveries count by POSTAL CODE and TIME WINDOW.
`

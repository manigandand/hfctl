# hfctl

# Recipe Stats Calculator

In the given assignment we suggest you to process an automatically generated JSON file with recipe data and calculated some stats.

## Instructions

1. Clone this repository.
2. Create a new branch called `dev`.
3. Create a pull request from your `dev` branch to the master branch.
4. Reply to the thread you're having with our HR department telling them we can start reviewing your code

## Given

Json fixtures file with recipe data. Download [Link](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz)

_Important notes_

1. Property value `"delivery"` always has the following format: "{weekday} {h}AM - {h}PM", i.e. "Monday 9AM - 5PM"
2. The number of distinct postcodes is lower than `1M`, one postcode is not longer than `10` chars.
3. The number of distinct recipe names is lower than `2K`, one recipe name is not longer than `100` chars.

## Functional Requirements

1. Count the number of unique recipe names.
2. Count the number of occurences for each unique recipe name (alphabetically ordered by recipe name).
3. Find the postcode with most delivered recipes.
4. Count the number of deliveries to postcode `10120` that lie within the delivery time between `10AM` and `3PM`, examples _(`12AM` denotes midnight)_:
   - `NO` - `9AM - 2PM`
   - `YES` - `10AM - 2PM`
5. List the recipe names (alphabetically ordered) that contain in their name one of the following words:
   - Potato
   - Veggie
   - Mushroom

## Non-functional Requirements

1. The application is packaged with [Docker](https://www.docker.com/).
2. Setup scripts are provided.
3. The submission is provided as a `CLI` application.
4. The expected output is rendered to `stdout`. Make sure to render only the final `json`. If you need to print additional info or debug, pipe it to `stderr`.
5. It should be possible to (implementation is up to you):  
   a. provide a custom fixtures file as input  
   b. provide custom recipe names to search by (functional reqs. 5)  
   c. provide custom postcode and time window for search (functional reqs. 4)

## Expected output

Generate a JSON file of the following format:

```json5
{
  unique_recipe_count: 15,
  count_per_recipe: [
    {
      recipe: "Mediterranean Baked Veggies",
      count: 1,
    },
    {
      recipe: "Speedy Steak Fajitas",
      count: 1,
    },
    {
      recipe: "Tex-Mex Tilapia",
      count: 3,
    },
  ],
  busiest_postcode: {
    postcode: "10120",
    delivery_count: 1000,
  },
  count_per_postcode_and_time: {
    postcode: "10120",
    from: "11AM",
    to: "3PM",
    delivery_count: 500,
  },
  match_by_name: [
    "Mediterranean Baked Veggies",
    "Speedy Steak Fajitas",
    "Tex-Mex Tilapia",
  ],
}
```

**General criteria from most important to less important**:

1. Functional and non-functional requirements are met.
2. Prefer application efficiency over code organisation complexity.
3. Code is readable and comprehensible. Setup instructions and run instructions are provided.
4. Command of the programming language that you chose.
5. Supporting notes on taken decisions and further clarifications are welcome.

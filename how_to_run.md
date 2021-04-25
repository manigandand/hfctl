## hfctl

CLI app - HelloFresh Recipe Stats Calculator

## How to run:

> Run via Docker (default test data)

```
./docker_run.sh
```

> Run via Docker with file input. 1st arg - absolute filepath

```
./docker_run.sh $PWD/test/hf_test_calculation_fixtures.json
```

> Run via Docker with file input and search cmd options.

- 1st arg - **absolute filepath**
- 2nd arg - **postcode**
- 3rd arg - **time window**
- 4th arg - **recipes to search**

```
./docker_run.sh $PWD/test/hf_test_calculation_fixtures.json "10121" "10AM - 2PM" "Mango"
```

---

> Run via binary

- 1st arg - **absolute filepath**
- 2nd arg - **postcode**
- 3rd arg - **time window**
- 4th arg - **recipes to search**

```
./deployment.sh
<!-- or -->
./deployment.sh $PWD/test/hf_test_calculation_fixtures.json
<!-- or -->
./deployment.sh $PWD/test/hf_test_calculation_fixtures.json "10121" "10AM - 2PM" "Mango" | jq .
```

---

## Known Issue

- memory & run time spikes up when loading large data sets

---

### Usage

`./hfctl -h`

```js
NAME:
   hfctl - CLI app to calculate HelloFresh Recipe Stats

USAGE:
   hfctl [global options] command [command options] [arguments...]
  eg:
  hfctl -f ./hf_fixtures.json search --recipes="Mango,Chicken" --postcode="10120" --time-window="10AM-2PM"


VERSION:
   v1.0.0

DESCRIPTION:
   Command Line Interface(CLI) for HelloFresh Recipe Stats Calculator.
   hfctl gives you a meaningful stats of recipes and let you search by postalcode and recipe names.

AUTHORS:
   HelloFresh
   Manigandan Dharmalingam <manigandan.jeff@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command
   Search Commands:
     search, s  Search recipes by name and count the no.of deliveries per postalcode between time window.

GLOBAL OPTIONS:
   --file FILE, -f FILE  Loads the recipe fixtures data from FILE
   --help, -h            show help (default: false)
   --version, -v         print the version (default: false)
```

---

`./hfctl search -h`

```js
Search recipes by name and count the no.of deliveries per postalcode between time window.

Examples:
  # Search all the recipes which has these RECIPE names.
  hfctl -f ./test/hf_test_calculation_fixtures.json search --recipes="Potato,Veggie,Mushroom"
  hfctl -f ./test/hf_test_calculation_fixtures.json search -r="Mango,Chicken"

  # Get the no.of deliveries count by POSTAL CODE and TIME WINDOW.
  hfctl -f ./test/hf_test_calculation_fixtures.json search --postcode="10120" --time-window="10AM-2PM"
  hfctl -f ./test/hf_test_calculation_fixtures.json search -pc="10120" -tw="10AM-2PM"

Options:
    -r, --recipes='': Search all the recipes which has these RECIPE KEY names.
    -pc, --postcode='': Get the no.of deliveries count by POSTAL CODE.
    -tw, --time-window='': Get the no.of deliveries count by POSTAL CODE and TIME WINDOW.
```

### Example CMD's:

`./hfctl --file test/hf_test_calculation_fixtures.json`

`./hfctl --file test/hf_test_calculation_fixtures.json search --postcode="10121" --time-window="10AM-2PM" --recipes="Mango,Chicken"`

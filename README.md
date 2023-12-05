# Go Advent of Code 2023

## Usage

### Running the code

To run the code for a given day you can youse `go run` command.
Make sure to provide the input file as the first argument and the part number as the second argument:

```sh
go run ./dayXY <input> <part> # where XY is the day number; part is either 1 or 2
```

For example:

```sh
go run ./day01 input.txt 1
```

This will print out the answer to the standard output.

### Tests

You can run tests, with the following command:

```sh
go test ./dayXY
```

The input for the tests is taken from the problem description.

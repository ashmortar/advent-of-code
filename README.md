# Advent of Code

This Repo contains the solutions and helpers I have created for doing the [Advent of Code](https://adventofcode.com/) challenges.

## Setup

- [install go](https://golang.org/doc/install)
- create a file at the top level of the repo called `aoc_cookie`
- paste the value of your session cookie from the advent of code site into the file

## Running the code

### Run a specific day

```bash
go run ./aoc 2020 1
# or
go run ./aoc -year 2020 -day 1
```

### Run a specific year

```bash
go run ./aoc 2020
# or
go run ./aoc -year 2020
```

### Run all years

```bash
go run ./aoc
```

## Generating Setup Code

### Create the next day chronologically

```bash
go run ./generate
```

### Create a specific day

```bash
go run ./generate 2020 1
# or
go run ./generate -year 2020 -day 1
```

## Fork and use for your own solutions?

If you want to use this repo as a starting point for your own solutions, you can fork it and then run the following command to remove all the solutions and generate the setup code for all the days.

### Reset all years

```bash
go run ./generate -reset
```

### Reset a year

```bash
go run ./generate -reset -year 2020
```

### Reset a day

```bash
go run ./generate -reset -year 2020 -day 1
```

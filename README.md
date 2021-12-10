# Advent Of Code 2021

My solutions in go

## Setup

To automatically fetch the input, add your AoC Cookie's session value (the hash after `session=`) in a file called `.token` at root of the project.

## Usage

### Build all solutions

Binaries will be output in `./bin/`.

`make`

### Generate template for current day

Generates `./dayX/dayX.go` and `./inputs/test/dayX.txt`.

`make day`

### Execute specific day over personal input

Fetches the personal input and writes to `./inputs/dayX.txt` if not already fetched.

`./bin/dayX`

### Execute specific day over test input

Requires the test input to be written in `./inputs/test/dayX.txt`.

`./bin/dayX -test`
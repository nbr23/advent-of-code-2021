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

Requires the test input to be written in `./inputs/test/dayX/1/input.txt` and the solutions to part 1 and 2 in `./inputs/test/dayX/1/result_p1.txt`,`./inputs/test/dayX/1/result_p2.txt`.

If more than 1 test inputs exist for a problem, increment the `Y`  in `./inputs/test/dayX/Y/` and reproduce the same pattern as described above.

Run tests on a defined day:

`make testday DAY=01`

For the current day:

`make testday`

### Execute tests over all days

`make testall`
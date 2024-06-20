# About
Study codes, tests and performance analysis.

It has 3 main folders:
- `analysis`: a benchmark tool
- `challenges`: code challenges and it's algorithm solutions
- `structures`: implementation of data structures to solve the challenges

## Run setup (install dependencies to make the project work) - only once

`go install github.com/go-task/task/v3/cmd/task@latest`

## Run a challenge file code

`task run -- filename` (e.g. `task run -- cards_slice`)

## Run all challenge files inside a folder

`task run -- folder_name` (e.g. `task run -- cards`)

## Run tests

`task test`

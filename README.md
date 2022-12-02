# AdventOfCode
Repository containing solutions (hopefully) to the Advent of Code brain teasers.

---

## AOC COOKIE
In order to use this repo you will need to obtain your Advent of Code Session Cookie.

This can be done by:

1. Logging in to [AOC](https://adventofcode.com/) with your github (or similar)

2. Open the browsers console (right clicking and then opening inspect)

3. `GET` any input on the page (or reload it) and look at the request headers
    - It will exist `Network` -> `Name [1]` -> `Headers` -> `Request Headers` -> `cookie` -> `session`

4. Make a note of this cookie
    - One recommendation is storing it in a simple text file at `~/.config/aoc/token`
    - Or as an environment variable
        - `export AOC_SESSION_COOKIE=...`

## Usage
Currently no release is created for this but it can be run in the following manner:
```
go run src/main.go -day 1 -year 2022 -aoc_cookie ...
```
- The `aoc_cookie` argument is optional and it will be searched for in the following order:
    - `-aoc_cookie` argument parsed
    - `AOC_SESSION_COOKIE` environment variable
    - `~/.config/aoc/token` file

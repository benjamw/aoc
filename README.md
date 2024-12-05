### Quick Note
I forked https://github.com/alexchao26/advent-of-code-go as it seemed a good start for what I want to focus on, which is improve my golang.
Thanks Alex :)

## Running Locally
### Installation
GOPATH should be set to [homedir]/go. In Windows, that is %USERPROFILE%/go, and in Linux, that is ~/go (or $HOME/go).

Then the project files should be hosted somewhere else other than in the GOPATH.

Run `go mod tidy` to get the mod file up to date and cleaned up.
You can run `go clean -modcache` if any errors block the run of tidy. Then run tidy again.

### Requirements
Go 1.18+ is required because [generics][generics] are used in some algorithms.

Use `go run main.go -part <1 or 2>` will be usable to run the actual inputs for that day.

Use `go test -run RegExpToMatchFunctionNames .` to run examples and unit tests via the `main_test.go` files.

## Scripts
Makefile should be fairly self-documenting. Alternatively you can run the binaries yourself via `go run` or `go build`.

`make help` prints a help message.

### Make skeleton files
```sh
for ((i=1; i<26; i++)); do
  make skeleton DAY=$i YEAR=2024
done
```

Note that skeletons use [embed][embed] and __will not compile__ without an `input.txt` file located in the same folder. Input files can be made via `make input`.
```sh
make skeleton DAY=5 YEAR=2020
make input DAY=5 YEAR=2020 AOC_SESSION_COOKIE=your_cookie
```

### Fetch inputs and write to input.txt files
Requires passing your cookie from AOC from either `-cookie` flag, or `AOC_SESSION_COOKIE` env variable.
```sh
make input DAY=1 YEAR=2020
```

[embed]: https://pkg.go.dev/embed
[generics]: https://go.dev/blog/intro-generics

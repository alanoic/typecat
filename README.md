# typecat 🐱
Typecat is a lightweight utility for word transformation and linguistic mapping into cat-friendly language (or lolcat). Y let computers rite nonsense in terminal when kat can tranzlate 4 u?

## Installation

Build the project using the provided `Makefile` or the `go build` command. The dictionary file is embedded into the binary at compile time using Go's `embed` package, so the resulting executable is self‑contained and does not need the `data/dictionary.yml` file at runtime.

```sh
# build with make (default target invokes `go build`)
make build

# or build manually
go build -o bin/typecat ./cmd/typecat
```

## Usage

## Sources
The dictionary used in this project is adopted from http://www.dribin.org/dave/lolspeak/ by Dave Dribin.
The actual file is provided by the [lolspeak](https://github.com/matthewhadley/lolspeak) project by Matthew Hadley.
# typecat 🐱

Typecat is a command-line utility that transforms text files into "lolcat" or "catspeak" language. Instead of letting your terminal spout boring prose, typecat translates your content into playful feline-speak:
> *correct grammar becomes creative chaos, proper nouns stay majestic, and your files get the meow treatment they deserve*.

## Installation

### From source (requires Go 1.20+)
```sh
go install github.com/alwek/typecat@latest
```

### Building locally
```sh
# build with make
make build

# or build manually
go build -o bin/typecat .
```

The dictionary file is embedded into the binary at compile time using Go's `embed` package, so the resulting executable is self‑contained and does not need the `data/dictionary.yml` file at runtime.

## Usage

### Basic usage
```sh
typecat <path-to-file>
```

Reads the file and prints the transformed content to stdout.

### Examples

```sh
# Transform any file
typecat README.md
```

### Example transformation

**Input:**
```
The quick brown fox jumps over the lazy dog.
This is a demonstration of the typecat utility.
```

**Output (transformed to lolcat):**
```
TEH QUICK BROWN FOX JUMPS OVAR TEH LAZY DAWG.
DIS AR TEH DEMONSTRASHUN OV TEH TYPECAT UTILITY.
```

## Sources
The dictionary used in this project is adopted from http://www.dribin.org/dave/lolspeak/ by Dave Dribin.
The actual dictionary file is from the [lolspeak](https://github.com/matthewhadley/lolspeak) project by Matthew Hadley.
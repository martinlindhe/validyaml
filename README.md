# About

Command line tool to validate and pretty-print YAML syntax of input
files.


## Installation

Windows and macOS binaries are available under [Releases](https://github.com/martinlindhe/validyaml/releases)

Or install from source:

    go get -u github.com/martinlindhe/validyaml


## Usage

Exit code will be 0 if file is good.

    $ validyaml file.yaml
    OK: file.yaml

    $ curl http://site.com/file.yaml | validyaml
    OK: -


## Pretty-print

    $ validyaml -p file.yaml

    list:
      one: 1
      two: 2


## License

Under [MIT](LICENSE)

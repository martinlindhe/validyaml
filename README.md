# About

Command line tool to validate YAML syntax of input files.

This tool simply exposes the super fast [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) to the command line.


# Installation

    go get -u github.com/martinlindhe/validyaml


# Usage

    validyaml file.yaml

    OK: file.yaml

To pretty print the result:

    validyaml -p file.yaml

    list:
      one: 1
      two: 2


# License

Under [MIT](LICENSE)

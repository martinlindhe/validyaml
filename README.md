# About

Command line tool to validate and pretty-print YAML syntax of input
files, taking advantage of [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2).


# Installation

    go get -u github.com/martinlindhe/validyaml


# Usage

Exit code will be 0 if file is good.

    validyaml file.yaml

    OK: file.yaml


# Pretty-print

    validyaml -p file.yaml

    list:
      one: 1
      two: 2


# License

Under [MIT](LICENSE)

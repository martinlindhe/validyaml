package main

import (
	"fmt"
	"io/ioutil"
	"os"

	termutil "github.com/andrew-d/go-termutil"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

var (
	inFile = kingpin.Arg("file", "YAML file.").String()
	pretty = kingpin.Flag("pretty", "Pretty print result.").Short('p').Bool()
	quiet  = kingpin.Flag("quiet", "Don't output on success.").Short('q').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, err := readPipeOrFile(*inFile)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	var f interface{}
	err = yaml.Unmarshal(data, &f)
	if err != nil {
		fmt.Println("ERROR:", *inFile, err)
		os.Exit(1)
	}

	if *pretty {
		b, err := yaml.Marshal(f)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Printf(string(b))
	} else {
		if !*quiet {
			fmt.Println("OK:", *inFile)
		}
	}
}

// readPipeOrFile reads from stdin if pipe exists, else from provided file
func readPipeOrFile(fileName string) ([]byte, error) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		return ioutil.ReadAll(os.Stdin)
	}
	if fileName == "" {
		return nil, fmt.Errorf("no piped data and no file provided")
	}
	return ioutil.ReadFile(fileName)
}

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func action(ctx *cli.Context) error {
	config, errors := makeConfig(ctx)
	if errors != nil {
		hasErr := false
		for _, err := range errors {
			if err.Error() != "" {
				fmt.Fprintf(os.Stderr, err.Error()+"\n")
				hasErr = true
			}
		}
		if hasErr {
			fmt.Fprintf(os.Stderr, "\n")
		}
		cli.ShowAppHelp(ctx)
	} else {
		dependency, errors := gatherDependency(config)
		output(config, dependency, errors)
	}
	return nil
}

func output(config *Config, dependency *Dependency, errors []error) {
	for _, err := range errors {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
	}
	switch config.Format {
	case "dot":
		//outputDot(config.Output, dependency)
		fmt.Print(dependency.graph.String())
	case "csv":
		outputCsv(config.Output, dependency)
	case "tsv":
		outputTsv(config.Output, dependency)
	case "json":
		outputJSON(config.Output, dependency)
	default:
		outputDefault(config.Output, dependency)
	}
}

func gatherDependency(config *Config) (*Dependency, []error) {
	var errors []error
	dependency := newDependency()
	for _, path := range config.Paths {
		deps, err := extract(path, config)
		if err != nil {
			errors = append(errors, err...)
		} else {
			dependency.concat(deps)
		}
	}
	return dependency, errors
}

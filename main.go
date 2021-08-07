package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type projectConfig struct {
	Name         string
	localPath    string
	RepoURL      string
	StaticAssets bool
}

func setupParseFlags(w io.Writer, args []string) (projectConfig, error) {
	conf := projectConfig{}
	genCom := flag.NewFlagSet("scaffold-gen", flag.ExitOnError)
	genCom.SetOutput(w)
	genCom.StringVar(&conf.Name, "n", "", "Project name")
	genCom.StringVar(&conf.localPath, "d", "", "Project location on disk")
	genCom.StringVar(&conf.RepoURL, "r", "", "Project remote repository URL")
	genCom.BoolVar(&conf.StaticAssets, "s", false, "Project will have static assets or not")
	err := genCom.Parse(args) // Exit on error with status 2 if there is error during parsing -> see line 20

	if genCom.NArg() != 0 {
		return conf, errors.New("Wrong command or no positional parameters expected.") // No additional parameters accepted except scaffold-gen
	}

	return conf, err
}

func validateConf(conf projectConfig) []error {
	var validateErrors []error
	if len(conf.Name) == 0 {
		validateErrors = append(validateErrors, errors.New("Project name cannot be empty."))
	}
	if len(conf.localPath) == 0 {
		validateErrors = append(validateErrors, errors.New("Project location cannot be empty."))
	}
	if len(conf.RepoURL) == 0 {
		validateErrors = append(validateErrors, errors.New("Project remote repository URL cannot be empty."))
	}
	return validateErrors
}

func generateScaffold(w io.Writer, conf projectConfig) {
	fmt.Fprintf(w, "Generating scaffold for project %s in %s\n", conf.Name, conf.localPath)
}

func main() {
	conf, err := setupParseFlags(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	errors := validateConf(conf)
	if len(errors) != 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(1)
	}

	generateScaffold(os.Stdout, conf)
}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var minimalPackage = []string{"doc", "main", "test"}

type gootFile struct {
	anchor, packName, fileName string
	okMessage, subcommand      string
	template                   Parseble
	output                     io.Writer
}

type filesList []gootFile

func (fl filesList) Process() error {
	for _, file := range fl {
		err := file.performCreation()
		if err != nil {
			return err
		}
	}
	return nil
}

// Checks if a given file is part of the mininmal
// version of the package.
func (gf gootFile) isMinimalFile() bool {
	for _, value := range minimalPackage {
		if gf.anchor == value {
			return true
		}
	}
	return false
}

// Checks if a given file should be created.
// This function is called is the user passess
// the --no-file_name subcommand
func (gf gootFile) shoudCreateFile() bool {
	subcNoPrefix := strings.TrimPrefix(strings.TrimPrefix(gf.subcommand, "--no"), "-")
	subcFiles := strings.Split(subcNoPrefix, "-")
	for _, file := range subcFiles {
		if gf.anchor == file {
			return false
		}
	}
	return true
}

// Creates the based on the construction
// passed on the gootstrap file
func (gf gootFile) createFile() error {

	// Creates the file and defer its closing
	fileCreate, err := os.Create(gf.fileName)
	if err != nil {
		return err
	}
	defer fileCreate.Close()

	// Writes the template into file and
	// then, writes the output to os.Stdout.
	fileCreate.WriteString(gf.template.Parse())
	fmt.Fprintln(gf.output, gf.okMessage)

	return nil
}

// Performs creation based on the subcommand passed
func (gf gootFile) performCreation() error {

	// Checks if the subcommand is either ""
	// or --minimal in order to perform the
	// creation on the correct files.
	// If defaults, it checks if the subcommand
	// matchs the --no-file_name pattern.
	switch gf.subcommand {
	case "":
		return createOrErrorOut(gf)
	case "--minimal":
		if gf.isMinimalFile() {
			return createOrErrorOut(gf)
		}
	default:
		// Checks if the subcommand passed is
		// related to removing files while
		// creating the package.
		if matchRemoveFile := findMatch("--no", gf.subcommand); matchRemoveFile != "" {
			if gf.shoudCreateFile() {
				return createOrErrorOut(gf)
			}
		}
	}

	return nil

}

func createOrErrorOut(gf gootFile) error {
	if err := gf.createFile(); err != nil {
		return err
	}
	return nil
}

//Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.
package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var knownSubcommands = []string{"--minimal"}

// Runs the program.
func run(args []string, out io.Writer) {
	switch len(args) {
	case 1:
		fmt.Fprintln(out, "===> Not enough arguments. Try goootstrap new project_name")
		return
	case 2:
		fmt.Fprintln(out, "===> You should set the name of your package. Try goootstrap new project_name\n")
		return
	default:
		runCommand(args, out)
	}
}

// Runs to program based on the command passed.
func runCommand(args []string, out io.Writer) {

	// Inits the command and the pack_name vars and
	// tests if there is any subcommand passed as
	// argument.
	command := args[1]
	pack_name := args[2]
	subcommand, isSubcKnown := func(args []string) (string, bool) {
		if len(args) > 3 {
			for _, value := range knownSubcommands {
				if args[3] == value {
					return args[3], true
				}
			}
			return args[3], false
		}
		return "", true
	}(args)

	switch command {
	case "new":
		// If the subcommand is known, it will
		// pass it along, if not, it will print an error message.
		if !isSubcKnown {
			fmt.Fprintf(out, "===> Subcommand %s unknown. Try typing one included in following list instead: %s\n", subcommand, strings.Join(knownSubcommands, ", "))
		} else {
			fmt.Fprintf(out, "===> Creating package %s\n", pack_name)
			createPackage(pack_name, subcommand, out)
			fmt.Fprintf(out, "===> Package created! cd %s to access.\n", pack_name)
		}
	default:
		fmt.Fprintf(out, "===> Command %s unknown. Try typing the command 'new' instead.\n", command)
	}
}

func main() {
	run(os.Args, os.Stdout)
}

// Creates the package with files in it
func createPackage(pack_name, subcommand string, out io.Writer) {
	sep := string(filepath.Separator)

	// Creates the project's folder

	if _, err := os.Stat(pack_name); os.IsNotExist(err) {
		os.Mkdir(pack_name, 0777)
		fmt.Fprintf(out, "===> Creating directory\n")
	}

	// Creates .gitignore

	gitignoreFile := gootFile{
		anchor:     "gitignore",
		fileName:   fmt.Sprintf("%s%s.gitignore", pack_name, sep),
		okMessage:  "===> Creating .gitignore file",
		output:     out,
		subcommand: subcommand,
	}
	err := gitignoreFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates .travis.yml

	travisFile := gootFile{
		anchor:     "travis",
		fileName:   fmt.Sprintf("%s%s.travis.yml", pack_name, sep),
		template:   travisTempl,
		okMessage:  "===> Creating .travis.yml file",
		output:     out,
		subcommand: subcommand,
	}
	err = travisFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates LISENCE.txt

	cuurentYear, _, _ := time.Now().Date()
	user, _ := user.Current()
	licenseFile := gootFile{
		anchor:     "license",
		fileName:   fmt.Sprintf("%s%sLICENSE.txt", pack_name, sep),
		template:   fmt.Sprintf(mitLicenseTempl, cuurentYear, user.Name),
		okMessage:  "===> Creating LICENSE.txt file",
		output:     out,
		subcommand: subcommand,
	}
	err = licenseFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates README.md

	readmeFile := gootFile{
		anchor:     "readme",
		fileName:   fmt.Sprintf("%s%sREADME.md", pack_name, sep),
		template:   fmt.Sprintf(readmeTempl, pack_name, pack_name),
		okMessage:  "===> Creating README.md file",
		output:     out,
		subcommand: subcommand,
	}
	err = readmeFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates main .go file

	mainFile := gootFile{
		anchor:     "main",
		fileName:   fmt.Sprintf("%s%s%s.go", pack_name, sep, pack_name),
		template:   mainTempl,
		okMessage:  fmt.Sprintf("===> Creating %s.go file", pack_name),
		output:     out,
		subcommand: subcommand,
	}
	err = mainFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates main _test.go file

	mainTestFile := gootFile{
		anchor:     "test",
		fileName:   fmt.Sprintf("%s%s%s_test.go", pack_name, sep, pack_name),
		template:   mainTestTempl,
		okMessage:  fmt.Sprintf("===> Creating %s_test.go file", pack_name),
		output:     out,
		subcommand: subcommand,
	}
	err = mainTestFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates doc.go file

	doctFile := gootFile{
		anchor:     "doc",
		fileName:   fmt.Sprintf("%s%sdoc.go", pack_name, sep),
		template:   docTempl,
		okMessage:  "===> Creating doc.go file",
		output:     out,
		subcommand: subcommand,
	}
	err = doctFile.performCreation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

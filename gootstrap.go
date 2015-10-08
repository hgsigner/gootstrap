//Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.
package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

//It runs the program.
func run(args []string, out io.Writer) {
	switch len(args) {
	case 1:
		fmt.Fprintln(out, "===> Not enough arguments. Try goootstrap new project_name")
		return
	case 2:
		fmt.Fprintln(out, "===> You should set the name of your package. Try goootstrap new project_name\n")
		return
	default:
		runCommand(args[1], args[2], out)
	}
}

//It runs to program based on the command passed.
func runCommand(command, pack_name string, out io.Writer) {
	switch command {
	case "new":
		fmt.Fprintf(out, "===> Creating package %s\n", pack_name)
		createPackage(pack_name, out)
		fmt.Fprintf(out, "===> Package created! cd %s to access.\n", pack_name)
	default:
		fmt.Fprintf(out, "===> Command %s unknown. Try typing the command 'new' instead.\n", command)
	}
}

func main() {
	run(os.Args, os.Stdout)
}

//It creates the package with files in it
func createPackage(pack_name string, out io.Writer) {
	sep := string(filepath.Separator)

	// Creates the project's folder

	if _, err := os.Stat(pack_name); os.IsNotExist(err) {
		os.Mkdir(pack_name, 0777)
		fmt.Fprintf(out, "===> Creating directory\n")
	}

	//Creates .gitignore

	gitignoreFile := gootFile{
		fileName:  fmt.Sprintf("%s%s.gitignore", pack_name, sep),
		okMessage: "===> Creating .gitignore file",
		output:    out,
	}
	err := gitignoreFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creates .travis.yml

	travisFile := gootFile{
		fileName:  fmt.Sprintf("%s%s.travis.yml", pack_name, sep),
		template:  travisTempl,
		okMessage: "===> Creating .travis.yml file",
		output:    out,
	}
	err = travisFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creates LISENCE.txt

	cuurentYear, _, _ := time.Now().Date()
	user, _ := user.Current()
	licenseFile := gootFile{
		fileName:  fmt.Sprintf("%s%sLICENSE.txt", pack_name, sep),
		template:  fmt.Sprintf(mitLicenseTempl, cuurentYear, user.Name),
		okMessage: "===> Creating LICENSE.txt file",
		output:    out,
	}
	err = licenseFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creates README.md

	readmeFile := gootFile{
		fileName:  fmt.Sprintf("%s%sREADME.md", pack_name, sep),
		template:  fmt.Sprintf(readmeTempl, pack_name, pack_name),
		okMessage: "===> Creating README.md file",
		output:    out,
	}
	err = readmeFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates main .go file

	mainFile := gootFile{
		fileName:  fmt.Sprintf("%s%s%s.go", pack_name, sep, pack_name),
		template:  mainTempl,
		okMessage: fmt.Sprintf("===> Creating %s.go file", pack_name),
		output:    out,
	}
	err = mainFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates main _test.go file

	mainTestFile := gootFile{
		fileName:  fmt.Sprintf("%s%s%s_test.go", pack_name, sep, pack_name),
		template:  mainTestTempl,
		okMessage: fmt.Sprintf("===> Creating %s_test.go file", pack_name),
		output:    out,
	}
	err = mainTestFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doctFile := gootFile{
		fileName:  fmt.Sprintf("%s%sdoc.go", pack_name, sep),
		template:  docTempl,
		okMessage: "===> Creating doc.go file",
		output:    out,
	}
	err = doctFile.createFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

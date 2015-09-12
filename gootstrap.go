//Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type gFile struct {
	packName, fileName, template, okMessage string
}

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

	gitignoreFile := gFile{
		fileName:  fmt.Sprintf("%s%s.gitignore", pack_name, sep),
		okMessage: "===> Creating .gitignore file\n",
	}
	err := createFile(gitignoreFile, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creates README.md

	readmeFile := gFile{
		fileName:  fmt.Sprintf("%s%sREADME.md", pack_name, sep),
		template:  fmt.Sprintf(readmeTempl, pack_name, pack_name),
		okMessage: "===> Creating README.md file\n",
	}
	err = createFile(readmeFile, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates main .go file

	mainFile := gFile{
		fileName:  fmt.Sprintf("%s%s%s.go", pack_name, sep, pack_name),
		template:  mainTempl,
		okMessage: fmt.Sprintf("===> Creating %s.go file\n", pack_name),
	}
	err = createFile(mainFile, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creates main _test.go file

	mainTestFile := gFile{
		fileName:  fmt.Sprintf("%s%s%s_test.go", pack_name, sep, pack_name),
		template:  mainTestTempl,
		okMessage: fmt.Sprintf("===> Creating %s_test.go file\n", pack_name),
	}
	err = createFile(mainTestFile, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doctFile := gFile{
		fileName:  fmt.Sprintf("%s%sdoc.go", pack_name, sep),
		template:  fmt.Sprintf(docTempl, pack_name),
		okMessage: "===> Creating doc.go file\n",
	}
	err = createFile(doctFile, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createFile(file gFile, out io.Writer) error {
	fileCreate, err := os.Create(file.fileName)

	if err != nil {
		return err
	}
	defer fileCreate.Close()

	if file.template != "" {
		fileCreate.WriteString(file.template)
	}

	fmt.Fprintln(out, file.okMessage)

	return nil
}

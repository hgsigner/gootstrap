package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func contains(t *testing.T, s, cont string) {
	if !strings.Contains(s, cont) {
		t.Fatalf("%s\nDoes not contain: %s", s, cont)
	}
}

func notContains(t *testing.T, s, cont string) {
	if strings.Contains(s, cont) {
		t.Fatalf("%s\nShould not contain: %s", s, cont)
	}
}

func Test_CreatePackageOk(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	contains(t, res, "===> Creating .gitignore file")
	contains(t, res, "===> Creating .travis.yml file")
	contains(t, res, "===> Creating README.md file")
	contains(t, res, "===> Creating LICENSE.txt file")
	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
	contains(t, res, "===> Creating doc.go file")
	contains(t, res, "===> Package created! cd new_package to access.")
}

func Test_CreateMinimalPackageOk(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--minimal"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	notContains(t, res, "===> Creating .gitignore file")
	notContains(t, res, "===> Creating .travis.yml file")
	notContains(t, res, "===> Creating README.md file")
	notContains(t, res, "===> Creating LICENSE.txt file")

	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
	contains(t, res, "===> Creating doc.go file")
	contains(t, res, "===> Package created! cd new_package to access.")
}

func Test_WithWrongSubcommand(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "balala"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	contains(t, res, "===> Subcommand balala unknown. Try typing one of the following: --minimal")
}

func Test_WithOneArg(t *testing.T) {

	command := []string{"gootstrap"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	contains(t, res, "===> Not enough arguments. Try goootstrap new project_name")
}

func Test_WithTwoArgs(t *testing.T) {

	command := []string{"gootstrap", "new"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	contains(t, res, "===> You should set the name of your package. Try goootstrap new project_name")
}

func Test_WithThreeArgsCommandNotOk(t *testing.T) {

	command := []string{"gootstrap", "fizz", "buzz"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	contains(t, res, "===> Command fizz unknown. Try typing the command 'new' instead.")
}

func Test_CreatePackageExcludingFileOk(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--no-gitignore-travis"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	notContains(t, res, "===> Creating .gitignore file")
	notContains(t, res, "===> Creating .travis.yml file")

	contains(t, res, "===> Creating README.md file")
	contains(t, res, "===> Creating LICENSE.txt file")
	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
	contains(t, res, "===> Creating doc.go file")
	contains(t, res, "===> Package created! cd new_package to access.")

}

func Test_CreatePackageExcludingFileNotOK(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--no"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	contains(t, res, "===> Creating .gitignore file")
	contains(t, res, "===> Creating .travis.yml file")
	contains(t, res, "===> Creating README.md file")
	contains(t, res, "===> Creating LICENSE.txt file")
	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
	contains(t, res, "===> Creating doc.go file")
	contains(t, res, "===> Package created! cd new_package to access.")

}

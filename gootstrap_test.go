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
	contains(t, res, "===> Creating CHANGELOG.md file")
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
	notContains(t, res, "===> Creating CHANGELOG.md file")

	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
	contains(t, res, "===> Creating doc.go file")
	contains(t, res, "===> Package created! cd new_package to access.")
}

func Test_CreateLightPackageOk(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--light"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	notContains(t, res, "===> Creating .gitignore file")
	notContains(t, res, "===> Creating .travis.yml file")
	notContains(t, res, "===> Creating README.md file")
	notContains(t, res, "===> Creating LICENSE.txt file")
	notContains(t, res, "===> Creating CHANGELOG.md file")
	notContains(t, res, "===> Creating doc.go file")

	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
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
	contains(t, res, "===> Creating CHANGELOG.md file")
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
	contains(t, res, "===> Creating CHANGELOG.md file")
	contains(t, res, "===> Creating new_package.go file")
	contains(t, res, "===> Creating new_package_test.go file")
	contains(t, res, "===> Creating doc.go file")
	contains(t, res, "===> Package created! cd new_package to access.")

}

func Test_CreatePackageWithCustomTemplate(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--template", "examples/simple.toml"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	contains(t, res, "===> Creating directory new_package")
	contains(t, res, "===> Creating directory new_package/utils")
	contains(t, res, "===> Creating directory new_package/labs")
	contains(t, res, "===> Creating new_package/utils/utils.go file")
	contains(t, res, "===> Creating new_package/utils/utils_test.go file")
	contains(t, res, "===> Creating new_package/labs/labs.go file")
	contains(t, res, "===> Creating new_package/labs/labs_test.go file")
	contains(t, res, "===> Creating new_package/.gitignore file")
	contains(t, res, "===> Creating new_package/README.md file")
	contains(t, res, "===> Creating new_package/main.go file")
	contains(t, res, "===> Package created! cd new_package to access.")

}

func Test_CreatePackageWithCustomTemplate_NoTemplate(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--template"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	contains(t, res, "===> You should pass the full path of the template file.")
	notContains(t, res, "===> Package created! cd new_package to access.")
}

func Test_CreatePackageWithCustomTemplate_LocalNotFound(t *testing.T) {
	command := []string{"gootstrap", "new", "new_package", "--template", "foobazbizz.toml"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	contains(t, res, "===> Error:")
	notContains(t, res, "===> Package created! cd new_package to access.")
}

func Test_CreatePackageWithCustomTemplate_Placeholder(t *testing.T) {

	command := []string{"gootstrap", "new", "new_package", "--template", "examples/placeholder.toml"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()
	contains(t, res, "===> Creating directory new_package")
	contains(t, res, "===> Creating new_package/new_package.go file")
	contains(t, res, "===> Creating new_package/new_package_test.go file")
	contains(t, res, "===> Creating new_package/README.md file")
	contains(t, res, "===> Package created! cd new_package to access.")
}

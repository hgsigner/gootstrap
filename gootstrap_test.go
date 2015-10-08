package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePackageOk(t *testing.T) {
	a := assert.New(t)

	command := []string{"gootstrap", "new", "new_package"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	a.Contains(res, "===> Creating .gitignore file")
	a.Contains(res, "===> Creating .travis.yml file")
	a.Contains(res, "===> Creating README.md file")
	a.Contains(res, "===> Creating LICENSE.txt file")
	a.Contains(res, "===> Creating new_package.go file")
	a.Contains(res, "===> Creating new_package_test.go file")
	a.Contains(res, "===> Creating doc.go file")
	a.Contains(res, "===> Package created! cd new_package to access.")
}

func Test_CreateMinimalPackageOk(t *testing.T) {
	a := assert.New(t)

	command := []string{"gootstrap", "new", "new_package", "--minimal"}

	w := &bytes.Buffer{}

	run(command, w)
	defer os.RemoveAll(command[2])

	res := w.String()

	a.NotContains(res, "===> Creating .gitignore file")
	a.NotContains(res, "===> Creating .travis.yml file")
	a.NotContains(res, "===> Creating README.md file")
	a.NotContains(res, "===> Creating LICENSE.txt file")

	a.Contains(res, "===> Creating new_package.go file")
	a.Contains(res, "===> Creating new_package_test.go file")
	a.Contains(res, "===> Creating doc.go file")
	a.Contains(res, "===> Package created! cd new_package to access.")
}

func Test_WithWrongSubcommand(t *testing.T) {
	a := assert.New(t)

	command := []string{"gootstrap", "new", "new_package", "balala"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	a.Contains(res, "===> Subcommand balala unknown. Try typing one included in following list instead: --minimal")
}

func Test_WithOneArg(t *testing.T) {
	a := assert.New(t)

	command := []string{"gootstrap"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	a.Contains(res, "===> Not enough arguments. Try goootstrap new project_name")
}

func Test_WithTwoArgs(t *testing.T) {
	a := assert.New(t)

	command := []string{"gootstrap", "new"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	a.Contains(res, "===> You should set the name of your package. Try goootstrap new project_name")
}

func Test_WithThreeArgsCommandNotOk(t *testing.T) {
	a := assert.New(t)

	command := []string{"gootstrap", "fizz", "buzz"}

	w := &bytes.Buffer{}
	run(command, w)
	res := w.String()

	a.Contains(res, "===> Command fizz unknown. Try typing the command 'new' instead.")
}

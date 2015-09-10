package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_name(t *testing.T) {
	a := assert.New(t)

	command := "new"
	pack_name := "new_package"

	w := &bytes.Buffer{}

	run(command, pack_name, w)
	//defer os.RemoveAll(pack_name)

	res := w.String()

	a.Contains(res, "===> Creating .gitignore file")
	a.Contains(res, "===> Creating README.md file")
	a.Contains(res, "===> Creating main .go file")
	a.Contains(res, "===> Creating doc.go file")
	a.Contains(res, "===> Package created! cd new_package to access.")
}

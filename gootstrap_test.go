package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_name(t *testing.T) {
	a := assert.New(t)

	command := "new"
	pack_name := "new_package"

	w := &bytes.Buffer{}

	run(command, pack_name, w)
	defer os.RemoveAll(pack_name)

	res := w.String()

	a.Contains(res, "===> Package created! cd new_package to access.")
}

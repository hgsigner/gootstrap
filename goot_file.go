package main

import (
	"fmt"
	"io"
	"os"
)

type gootFile struct {
	packName, fileName, template, okMessage string
	output                                  io.Writer
}

func (gf gootFile) createFile() error {
	fileCreate, err := os.Create(gf.fileName)

	if err != nil {
		return err
	}
	defer fileCreate.Close()

	if gf.template != "" {
		fileCreate.WriteString(gf.template)
	}

	fmt.Fprintln(gf.output, gf.okMessage)

	return nil
}

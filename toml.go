package main

import (
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type tomlConfig struct {
	Directories []directory
	Files       []file
}

type file struct {
	Name, Template, Dir string
}

type directory struct {
	Name  string
	Files []file
}

func NewTomlTemplate(filepath string) (*tomlConfig, error) {

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var config tomlConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

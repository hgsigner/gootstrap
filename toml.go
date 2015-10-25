package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

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

// Creates a new instance of the tomlConfig
// struct and reads the file accordingly, checkins
// its a local a file or a remote file.
func NewTomlTemplate(fileLocation string) (*tomlConfig, error) {

	var buf []byte
	var err error

	reg := regexp.MustCompile("^(http|https)://")
	if reg.MatchString(fileLocation) {
		fmt.Fprintf(os.Stdout, "===> Fetiching url: %s\n", fileLocation)
		buf, err = getRemoteFile(fileLocation)
		if err != nil {
			return nil, err
		}
	} else {
		buf, err = getLocalFile(fileLocation)
		if err != nil {
			return nil, err
		}
	}

	var config tomlConfig
	if err = toml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Reads local files
func getLocalFile(filePath string) ([]byte, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Reads remote files
func getRemoteFile(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("File not found on remote. Check your url again.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

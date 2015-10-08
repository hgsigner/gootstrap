package main

//README.md template
var readmeTempl = `#%s

This is the awesome description for %s.

##Licensing
Add some licensing to your package`

//doc.go template
var docTempl = `// Add some documentation to your package.
package %s`

//Main .go file template
var mainTempl = `package main

import (
	"fmt"
)

func main() {
	fmt.Prinln("Hello from Gootstrap!")
}
`

//Main _test.go file template
var mainTestTempl = `package main

import (
	"testing"
)

func Test(t *testing.T) {
	
}
`

var travisTempl = `language: go
sudo: false

go:
  - 1.3
  - 1.4
  - 1.5
  - tip

script:
  - go test -v ./...
`

package main

//README.md template
var readmeTempl = `#%s

This is the awesome description for %s.`

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

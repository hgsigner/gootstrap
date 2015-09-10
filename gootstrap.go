package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var readmeTempl = `#%s
This is the awesome description for %s.`

var docTempl = `// Add some documentation to your package.
package %s`

var mainTempl = `package main
import (
	"fmt"
)

func main() {
	fmt.Prinln("Hello from Gootstrap!")
}
`

func main() {

	pack_name := os.Args[1]
	fmt.Println(pack_name)

	sep := string(filepath.Separator)

	// Creates the project's folder
	err := os.Mkdir(pack_name, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Creates README.md
	readme := fmt.Sprintf("%s%sREADME.md", pack_name, sep)
	readme_file, err := os.Create(readme)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer readme_file.Close()
	fReadme := fmt.Sprintf(readmeTempl, pack_name, pack_name)
	readme_file.WriteString(fReadme)

	// Creates main .go file
	mainpack := fmt.Sprintf("%s%s%s.go", pack_name, sep, pack_name)
	mainpack_file, err := os.Create(mainpack)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer mainpack_file.Close()
	mainpack_file.WriteString(mainTempl)

	// Creates main doc.go file
	doc := fmt.Sprintf("%s%sdoc.go", pack_name, sep)
	doc_file, err := os.Create(doc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer doc_file.Close()
	dReadme := fmt.Sprintf(docTempl, pack_name)
	doc_file.WriteString(dReadme)

}

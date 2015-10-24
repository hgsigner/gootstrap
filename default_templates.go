package main

var changeLogTmpl = `#Changelog

##0.0.1 - {{.Date}}

- Add some changelog to this version
`

var docTmpl = `// Add some documentation to your package.
package {{.PackName}}
`

var gitIgTmpl = `.DS_Store`

var licenseTmpl = `The MIT License (MIT)

Copyright (c) {{.Year}} {{.UserName}}

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
`

var mainTmpl = `package {{.PackName}}`

var readmeTmpl = `#{{.Title}}

This is the awesome description for {{.Project}}.

## License

This package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
`

var testTmpl = `package {{.PackName}}

import (
	"testing"
)

func Test(t *testing.T) {
	
}
`

var travisTmpl = `language: go
sudo: false

go:
  - 1.3
  - 1.4
  - 1.5
  - tip

script:
  - go test -v ./...
`

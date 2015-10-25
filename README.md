#Gootstrap [![Build Status](https://travis-ci.org/hgsigner/gonumbers.svg?branch=master)](https://travis-ci.org/hgsigner/gonumbers)

Gootstrap is a simple package that bootstraps new Go packages. I've created it because I was repeating myself a lot while starting new projects.

##Installing:

```bash
$ go get github.com/hgsigner/gootstrap
```

##Usage:

###Default Package:

After installing, you can use **gootstrap new package_name** to create a new project.

```bash
$ gootstrap new project_name
===> Creating package project_name
===> Creating directory
===> Creating .gitignore file
===> Creating .travis.yml file
===> Creating README.md
===> Creating LICENSE.txt file
===> Creating CHANGELOG.md file
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name

|-- project_name
    |-- .gitignore
    |-- .travis.yml
    |-- README.md
    |-- LICENSE.txt
    |-- CHANGELOG.md
    |-- project_name.go
    |-- project_name_test.go
    |-- doc.go
```

If the command gootstrap does not work for you, use `$GOPATH/bin/gootstrap new project_name` instead.

####--minimal:

In order to create a minimal package structure **(main file, test file and doc file)**, pass the `--minimal` flag after the package name:

```bash
$ gootstrap new project_name --minimal
===> Creating package project_name
===> Creating directory
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name

|-- project_name
    |-- project_name.go
    |-- project_name_test.go
    |-- doc.go
```

####--light:

In order to create a light package structure **(main file and test file)**, pass the `--light` flag after the package name:

```bash
$ gootstrap new project_name --light
===> Creating package project_name
===> Creating directory
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Package created! cd project_name to access.
$ cd project_name

|-- project_name
    |-- project_name.go
    |-- project_name_test.go
```

####--no-{file names}:

If you want the exclude some files while creating the package, you can pass the subcommand `--no-{file names separated by "-"}`:

```bash
$ gootstrap new project_name --no-travis-license
===> Creating package project_name
===> Creating directory
===> Creating .gitignore file
===> Creating README.md
===> Creating CHANGELOG.md file
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name

|-- project_name
    |-- .gitignore
    |-- REAMDE.md
    |-- CHANGELOG.md
    |-- project_name.go
    |-- project_name_test.go
    |-- doc.go
```

File names you can pass to `--no-{files separated by "-"}`: `travis, gitignore, license, readme, main, test, doc and changelog`

###Custom Templates:

Gootstrap allows you to create your own template file and use it as reference when creating your package. Gootstrap uses `*.toml` files in order organize the templates.

Lets assume that you have the following `example.toml` file in your file's system:

```tmol
#Creates the directories

[[directories]]
name = "utils"
	[[directories.files]]
	name = "utils.go"
	template = '''package utils'''
	[[directories.files]]
	name = "utils_test.go"
	template = '''package utils

	import "testing"

	func Test(t *testing.T) {

	}
	'''

[[directories]]
name = "labs"
	[[directories.files]]
	name = "labs.go"
	template = '''package labs'''
	
	[[directories.files]]
	name = "labs_test.go"
	template = '''package labs

	import "testing"

	func Test(t *testing.T) {

	}
	'''

#Creates files in the root directory

[[files]]
name = "README.md"
template = '''#Readme
some reamde
'''

[[files]]
name = "main.go"
template = '''package main

import "fmt"

func main() {
	fmt.Prinln("Hello!")
}
'''

```

Passing the **FULLPATH** of your template file as argument for `--template` flag on gootstrap, will produce the following result:

```bash
$ gootstrap new new_project --template /full/path/of/your/example.toml 
===> Creating package new_project
===> Creating directory new_project
===> Creating directory new_project/utils
===> Creating directory new_project/labs
===> Creating new_project/utils/utils.go file
===> Creating new_project/utils/utils_test.go file
===> Creating new_project/labs/labs.go file
===> Creating new_project/labs/labs_test.go file
===> Creating new_project/README.md file
===> Creating new_project/main.go file
===> Package created! cd new_project to access.
```

##Creating your own template:

Gootstrap only accepts the following arrays of tables as the structure for the templates:

```toml
[[directories]] # Creates directory
name
[[directories.files]] # Creates files inside the directory
name
template
[[files]] # Creates files inside the root directory
name
template
```

Now, lets create our custom template:

```toml
# example2.toml
# Creates directories inside the root folder
# and creates files inside this directory.
[[directories]]
# name of the directory
name = "routes"
	[[directories.files]]
	# Name of the file
	name = "routes.go"
	# Template for the file
	template = '''package routes
	
	import "fmt"
	
	func sayHi() {
		fmt.Println("Hi!")
	}
	'''
	[[directories.files]]
	# Name of the file
	name = "routes_test.go"
	# Template for the file
	template = '''package routes
	
	import "testing"
	
	func Test(t *testing.T) {
	
	}
	'''

#Creates files inside the root folder
[[files]]
name = "CHANGELOG.md"
template = '''#Changelog
Some changelog
'''

[[files]]
name = "main.go"
template = '''package main

import "fmt"

func main() {
	fmt.Prinln("Hello!")
}
'''
```

Lets use it:

```bash
$ gootstrap new new_project --template /full/path/of/your/example2.toml 
===> Creating package new_project_ex2
===> Creating directory new_project_ex2
===> Creating directory new_project_ex2/routes
===> Creating new_project_ex2/routes/routes.go file
===> Creating new_project_ex2/routes/routes_test.go file
===> Creating new_project_ex2/CHANGELOG.md file
===> Creating new_project_ex2/main.go file
===> Package created! cd new_project_ex2 to access.
```

###Placeholders:

Gootstrap allows you to insert placeholders inside your templates in order to replace it with custom text. The built in placeholders are:

- `{{.PackageName}}`: Gets the package name (gootstrap new **new_pack**);
- `{{.HumanizedPackageName}}`: Returns the humanized package name (e.g. **NewPackage**);
- `{{.CurrentYear}}`: Gets the current year;
- `{{.UserName}}`: Gets the user's computer name;
- `{{.Date}}`: Gets the current date (YYYY-MM-DD).

####Usage:

```toml
# placeholder.toml
[[files]]
name = "{{.PackageName}}.go"
template = '''package {{.PackageName}}'''
[[files]]
name = "{{.PackageName}}_test.go"
template = '''package {{.PackageName}}

import "testing"

func Test(t *testing.T) {

}
'''
```

####Performing:

```bash
$ gootstrap new place_holder --template /full/path/of/your/placeholder.toml 
===> Creating package place_holder
===> Creating place_holder/place_holder.go file
===> Creating place_holder/place_holder_test.go file
===> Package created! cd place_holder to access.
```

####Remote Templates:

Gootstrap also allows you to fetch templates from remote servers, passing its url as the argument for the `--template` flag:

```bash
$ gootstrap new remote_simple --template https://raw.githubusercontent.com/hgsigner/gootstrap/master/examples/simple.toml
===> Creating package remote_simple
===> Fetiching url: https://raw.githubusercontent.com/hgsigner/gootstrap/master/examples/simple.toml
===> Creating directory remote_simple
===> Creating directory remote_simple/utils
===> Creating directory remote_simple/labs
===> Creating remote_simple/utils/utils.go file
===> Creating remote_simple/utils/utils_test.go file
===> Creating remote_simple/labs/labs.go file
===> Creating remote_simple/labs/labs_test.go file
===> Creating remote_simple/.gitignore file
===> Creating remote_simple/README.md file
===> Creating remote_simple/main.go file
===> Package created! cd remote_simple to access.
```

Creating your own template is simple and can save you a lot of time if you keep doing the same thing over and over again.
- - -
For more info on TOML, be sure to check [https://github.com/toml-lang/toml](https://github.com/toml-lang/toml)
- - -

Any bug or feedback, feel free to drop me a line :)

##Licensing
This package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).

#Gootstrap [![Build Status](https://travis-ci.org/hgsigner/gonumbers.svg?branch=master)](https://travis-ci.org/hgsigner/gonumbers)

Gootstrap is a simple package that bootstraps new Go packages. I've created it because I was repeating myself a lot while starting new projects.

##Installing

```bash
$ go get github.com/hgsigner/gootstrap
```

##Usage
After installing, you can use **gootstrap new package_name** to create a new project.

```bash
$ goostrap new project_name
===> Creating package project_name
===> Creating directory
===> Creating .gitignore file
===> Creating .travis.yml file
===> Creating README.md
===> Creating LICENSE.txt file
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name
```

If the command goostrap does not work for you, use "$GOPATH/bin/goostrap new project_name" instead.

###Tree
```
|-- project_name
    |-- .gitignore
    |-- .travis.yml
    |-- REAMDE.md
    |-- LICENSE.txt
    |-- project_name.go
    |-- project_name_test.go
    |-- doc.go
```

In order to create a minimal package structure, pass the `--minimal` argument after the package name:

```bash
$ goostrap new project_name --minimal
===> Creating package project_name
===> Creating directory
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name
```

###Tree
```
|-- project_name
    |-- project_name.go
    |-- project_name_test.go
    |-- doc.go
```

If you want the exclude some files while creating the package, you can pass the subcommand `--no-{file names separated by "-"}`:

```bash
$ goostrap new project_name --no-travis-license
===> Creating package project_name
===> Creating directory
===> Creating .gitignore file
===> Creating README.md
===> Creating project_name.go file
===> Creating project_name_test.go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name
```

###Tree
```
|-- project_name
    |-- .gitignore
    |-- REAMDE.md
    |-- project_name.go
    |-- project_name_test.go
    |-- doc.go
```

File names you can pass to `--no-{files separated by "-"}`: `travis, gitignore, license, readme, main, test and doc`

- - -
I know that there are tons of things I could do to make it better, but, for now, it meets my needs. I'll be improving it.
- - -
##Licensing
This package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
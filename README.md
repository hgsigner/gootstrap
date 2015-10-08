#Gootstrap

Gootstrap is a simple package that bootstraps new Go packages. I've created it because I was repeating myself a lot when starting new projects.

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

In order to create a minimal package structure, pass the `--minimal` arguments after the package name:

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

- - -
I know that there are tons of things I could do to make it better, but, for now, it meets my needs. I'll be improving it.
- - -
##Licensing
You can use or modify it to meet your needs.
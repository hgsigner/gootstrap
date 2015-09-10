#Gootstrap is a simple Go package bootstrapper.

Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.

##Installing

```bash
$ go get github.com/hgsigner/gootstrap
```

##Usage
After installing, you can use **gootstrap new package_name** to create a new project.

```bash
$ $GOPATH/bin/goostrap new project_name
2015/09/10 00:06:36 Creating package project_name
2015/09/10 00:06:36 Creating directory
2015/09/10 00:06:36 Creating README.md
2015/09/10 00:06:36 Creating main .go file.
2015/09/10 00:06:36 Creating doc.go
2015/09/10 00:06:36 Package created! cd project_name to access.
$ cd project_name
```
###Tree
```
|-- project_name
    |-- REAMDE.md
    |-- project_name.go
    |-- doc.go
```

I know that there is tons of things I could do to make it better, but, for now, it meets my needs. I be improving it.
- - -
##Licensing
You can use or modify it to meet your needs.
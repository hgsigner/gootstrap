#Gootstrap

Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.

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
===> Creating README.md
===> Creating main .go file
===> Creating doc.go file
===> Package created! cd project_name to access.
$ cd project_name
```

If the command goostrap does not work for you, use "$GOPATH/bin/goostrap new project_name" instead.

###Tree
```
|-- project_name
    |-- .gitignore
    |-- REAMDE.md
    |-- project_name.go
    |-- doc.go
```

I know that there are tons of things I could do to make it better, but, for now, it meets my needs. I be improving it.
- - -
##Licensing
You can use or modify it to meet your needs.
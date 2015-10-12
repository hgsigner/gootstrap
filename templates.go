package main

import (
	"bytes"
	"text/template"
)

var t *template.Template

func init() {
	t, _ = template.ParseGlob("./templates/*.tmpl")
}

type Parseble interface {
	Parse() string
}

//.gitignore template

type GitIgnoreFile struct {
}

func (gi GitIgnoreFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "gitignore.tmpl", nil)
	return w.String()
}

//README.md template

type ReadmeFile struct {
	Title   string
	Project string
}

func (rdm ReadmeFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "readme.tmpl", rdm)
	return w.String()
}

//doc.go template

type DocFile struct {
	PackName string
}

func (doc DocFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "doc.tmpl", doc)
	return w.String()
}

//Main .go file template

type MainFile struct {
	PackName string
}

func (m MainFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "main.tmpl", m)
	return w.String()
}

//Main _test.go file template

type MainTestFile struct {
	PackName string
}

func (mt MainTestFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "test.tmpl", mt)
	return w.String()
}

//Main .travis.yml file template

type TravisFile struct {
}

func (tv TravisFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "travis.tmpl", nil)
	return w.String()
}

//Main LICENSE.txt file template

type LicenseFile struct {
	Year     int
	UserName string
}

func (ls LicenseFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "license.tmpl", ls)
	return w.String()
}

//Main CHANGELOG.md file template

type ChangelogFile struct {
	Date string
}

func (cl ChangelogFile) Parse() string {
	w := &bytes.Buffer{}
	t.ExecuteTemplate(w, "changelog.tmpl", cl)
	return w.String()
}

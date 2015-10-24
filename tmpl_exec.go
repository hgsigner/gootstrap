package main

import (
	"bytes"
	"text/template"
)

type Parseble interface {
	Parse() string
}

//.gitignore parser

type GitIgnoreFile struct {
}

func (gi GitIgnoreFile) Parse() string {
	return parse_templates("gitignore", gitIgTmpl, gi)
}

//README.md parser

type ReadmeFile struct {
	Title   string
	Project string
}

func (rdm ReadmeFile) Parse() string {
	return parse_templates("readme", readmeTmpl, rdm)
}

//doc.go parser

type DocFile struct {
	PackName string
}

func (doc DocFile) Parse() string {
	return parse_templates("doc", docTmpl, doc)
}

//Main .go file parser

type MainFile struct {
	PackName string
}

func (m MainFile) Parse() string {
	return parse_templates("main", mainTmpl, m)
}

//Main _test.go file parser

type MainTestFile struct {
	PackName string
}

func (mt MainTestFile) Parse() string {
	return parse_templates("test", testTmpl, mt)
}

//Main .travis.yml file parser

type TravisFile struct {
}

func (tv TravisFile) Parse() string {
	return parse_templates("travis", travisTmpl, tv)
}

//Main LICENSE.txt file parser

type LicenseFile struct {
	Year     int
	UserName string
}

func (ls LicenseFile) Parse() string {
	return parse_templates("license", licenseTmpl, ls)
}

//Main CHANGELOG.md file parser

type ChangelogFile struct {
	Date string
}

func (cl ChangelogFile) Parse() string {
	return parse_templates("changelog", changeLogTmpl, cl)
}

// Custom template parser

type CustomTemplate struct {
	PackageName          string
	HumanizedPackageName string
	CurrentYear          int
	UserName             string
	Date                 string
	Template             string
}

func (ct CustomTemplate) Parse() string {
	return parse_templates("customtemplate", ct.Template, ct)
}

// Parse helper
func parse_templates(name, tmpl string, prs Parseble) string {
	w := &bytes.Buffer{}
	t := template.Must(template.New(name).Parse(tmpl))
	t.Execute(w, prs)
	return w.String()
}

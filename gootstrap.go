//Gootstrap is a simple package that bootstrap new Go packages. It creates a REAME.md, a doc.go and main (package name).go file as a placeholder. I did it because I was repeating myself a lot when starting new projects.
package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/hgsigner/stringfy"
)

var knownSubcommands = []string{"--minimal", "--no", "--template"}

// Runs the program.
func run(args []string, out io.Writer) {
	switch len(args) {
	case 1:
		fmt.Fprintln(out, "===> Not enough arguments. Try goootstrap new project_name")
		return
	case 2:
		fmt.Fprintln(out, "===> You should set the name of your package. Try goootstrap new project_name")
		return
	default:
		runCommand(args, out)
	}
}

// Runs to program based on the command passed.
func runCommand(args []string, out io.Writer) {

	// Inits the command and the pack_name vars and
	// tests if there is any subcommand passed as
	// argument.
	command := args[1]
	pack_name := args[2]
	subcommand, isSubcKnown := func(args []string) (string, bool) {
		if len(args) > 3 {
			for _, value := range knownSubcommands {
				v := findMatch(value, args[3])
				if v != "" {
					return args[3], true
				}
			}
			return args[3], false
		}
		return "", true
	}(args)

	switch command {
	case "new":
		// If the subcommand is known, it will
		// pass it along, if not, it will print an error message.
		if !isSubcKnown {
			fmt.Fprintf(out, "===> Subcommand %s unknown. Try typing one of the following: %s\n", subcommand, strings.Join(knownSubcommands, ", "))
		} else {
			fmt.Fprintf(out, "===> Creating package %s\n", pack_name)

			if subcommand != "--template" {
				createDefaultPackage(pack_name, subcommand, out)
			} else {
				// Checks if the template path was passed
				if len(args) < 5 {
					fmt.Fprintf(out, "===> You should pass the full path of the template file.\n")
					return
				}

				// Everything is ok.
				// Should create the package.
				createTemplatePackage(pack_name, args[4], out)
			}

			fmt.Fprintf(out, "===> Package created! cd %s to access.\n", pack_name)
		}
	default:
		fmt.Fprintf(out, "===> Command %s unknown. Try typing the command 'new' instead.\n", command)
	}
}

func main() {
	run(os.Args, os.Stdout)
}

// Creates the package with files in it
func createDefaultPackage(pack_name, subcommand string, out io.Writer) {
	sep := string(filepath.Separator)

	// Creates the project's folder
	createFolder(pack_name, out)

	// Init files

	currentYear, currentMonth, currentDay := time.Now().Date()
	currentDate := fmt.Sprintf("%d-%d-%d", currentYear, currentMonth, currentDay)
	user, _ := user.Current()

	files := filesList{
		{
			anchor:     "gitignore",
			fileName:   fmt.Sprintf("%s%s.gitignore", pack_name, sep),
			template:   GitIgnoreFile{},
			okMessage:  "===> Creating .gitignore file",
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:     "travis",
			fileName:   fmt.Sprintf("%s%s.travis.yml", pack_name, sep),
			template:   TravisFile{},
			okMessage:  "===> Creating .travis.yml file",
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:     "license",
			fileName:   fmt.Sprintf("%s%sLICENSE.txt", pack_name, sep),
			template:   LicenseFile{currentYear, user.Name},
			okMessage:  "===> Creating LICENSE.txt file",
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:     "readme",
			fileName:   fmt.Sprintf("%s%sREADME.md", pack_name, sep),
			template:   ReadmeFile{stringfy.CamelCase(pack_name), pack_name},
			okMessage:  "===> Creating README.md file",
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:     "main",
			fileName:   fmt.Sprintf("%s%s%s.go", pack_name, sep, pack_name),
			template:   MainFile{pack_name},
			okMessage:  fmt.Sprintf("===> Creating %s.go file", pack_name),
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:     "test",
			fileName:   fmt.Sprintf("%s%s%s_test.go", pack_name, sep, pack_name),
			template:   MainTestFile{pack_name},
			okMessage:  fmt.Sprintf("===> Creating %s_test.go file", pack_name),
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:   "doc",
			fileName: fmt.Sprintf("%s%sdoc.go", pack_name, sep),

			template:   DocFile{pack_name},
			okMessage:  "===> Creating doc.go file",
			output:     out,
			subcommand: subcommand,
		},
		{
			anchor:   "changelog",
			fileName: fmt.Sprintf("%s%sCHANGELOG.md", pack_name, sep),

			template:   ChangelogFile{currentDate},
			okMessage:  "===> Creating CHANGELOG.md file",
			output:     out,
			subcommand: subcommand,
		},
	}

	err := files.Process()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Creates the package with files in it
func createTemplatePackage(packName, templPath string, out io.Writer) {
	sep := string(filepath.Separator)

	// Inits a new instance of the toml parsed template
	tomlTempl, err := NewTomlTemplate(templPath)
	if err != nil {
		fmt.Fprintf(out, "===> Error: %s\n", err)
		os.Exit(1)
	}

	// Creates the project's folder
	createFolder(packName, out)

	// Loops through the template and creates the
	// folders and files for the folders
	files := make(filesList, 0)
	for _, dir := range tomlTempl.Directories {
		createFolder(packName+sep+dir.Name, out)
		for _, fl := range dir.Files {
			filepath := packName + sep + dir.Name + sep
			gf := createCustomGootFile(packName, fl.Name, fl.Template, filepath, out)
			files = append(files, gf)
		}
	}

	// Creates files in the root directory
	for _, fl := range tomlTempl.Files {
		filepath := packName + sep
		gf := createCustomGootFile(packName, fl.Name, fl.Template, filepath, out)
		files = append(files, gf)
	}

	// Processes the files
	err = files.Process()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// Creates the folders.
// Its a helper function.
func createFolder(folderName string, out io.Writer) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, 0777)
		fmt.Fprintf(out, "===> Creating directory %s\n", folderName)
	}
}

// Creates the custom gootFiles.
// Its a helper function.
func createCustomGootFile(packName, fileName, template, filepath string, out io.Writer) gootFile {
	currentYear, currentMonth, currentDay := time.Now().Date()
	currentDate := fmt.Sprintf("%d-%d-%d", currentYear, currentMonth, currentDay)
	user, _ := user.Current()

	fileNameTempl := CustomTemplate{PackageName: packName, Template: fileName}

	filename := filepath + fileNameTempl.Parse()

	gf := gootFile{
		fileName: filename,
		template: CustomTemplate{
			PackageName:          packName,
			HumanizedPackageName: stringfy.CamelCase(packName),
			CurrentYear:          currentYear,
			UserName:             user.Name,
			Date:                 currentDate,
			Template:             template,
		},
		okMessage: fmt.Sprintf("===> Creating %s file", filename),
		output:    out,
	}

	return gf
}

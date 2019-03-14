package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const startSize = 8

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Too few arguments...\n")
		return
	}

	// Take the first argument as path to clang template
	//pathTemplate := os.Args[1]

	// Take the second argument as path to CMakeLists.txt
	pathListFile := os.Args[1]

	// Read files into strings
	contentList, err := ioutil.ReadFile(pathListFile)
	if err != nil {
		panic(err)
	}
	//contentTemplate, err := ioutil.ReadFile(pathTemplate)
	//if err != nil {
	//	panic(err)
	//}

	// Find library or executable names
	targets := findLibraryNames(string(contentList))
	for i, match := range targets {
		fmt.Printf("match %d: %s\n", i, match)
	}

	// Find project name
	projectName, err := findProjectName(string(contentList))
	fmt.Printf("Project name: %s\n", projectName)
}

// findLibraryNames finds the names of libraries and executables defined
// by add_library and add_executable
func findLibraryNames(text string) []string {
	libNames := make([]string, 0, startSize)
	targetMatch := string(` *add_(?:library|executable)\( *(\w*)`)
	r := regexp.MustCompile(targetMatch)
	matches := r.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		libNames = append(libNames, match[1])
	}
	return libNames
}

// findLibraryNames finds the name of project in CMakeLists.txt
func findProjectName(text string) (string, error) {
	targetMatch := string(` *project\((\w*)\)`)
	r := regexp.MustCompile(targetMatch)
	matches := r.FindAllStringSubmatch(text, -1)
	if len(matches) > 1 {
		return "", errors.New("more than one project names were found")
	} else if len(matches) < 1 {
		return "", errors.New("no project name was found")
	}
	return matches[0][1], nil
}

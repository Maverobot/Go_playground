package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const startSize = 8

func main() {

	// Take the first argument as path to CMakeLists.txt
	var path string
	if len(os.Args) < 2 {
		fmt.Printf("Too few arguments...\n")
		return
	}
	path = os.Args[1]

	// Read file into a string
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Find library or executable names
	matches := findLibraryNames(string(content))
	for i, match := range matches {
		fmt.Printf("match %d: %s\n", i, match)
	}
}

// findLibraryNames find the names of libraries and executables defined
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

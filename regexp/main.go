package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const startSize = 8

const template = `
## ClangTools
include(${CMAKE_CURRENT_LIST_DIR}/../cmake/ClangTools.cmake OPTIONAL
  RESULT_VARIABLE CLANG_TOOLS
)
if(CLANG_TOOLS)
  file(GLOB_RECURSE SOURCES
    ${CMAKE_CURRENT_SOURCE_DIR}/src/*.cpp)
  file(GLOB_RECURSE HEADERS
    ${CMAKE_CURRENT_SOURCE_DIR}/include/*.h
    ${CMAKE_CURRENT_SOURCE_DIR}/src/*.h
  )
  add_format_target(${PROJECT_NAME} FILES ${SOURCES} ${HEADERS})
  add_tidy_target(${PROJECT_NAME}
    FILES ${SOURCES}
    DEPENDS ${TARGETS}
  )
endif()
`

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Too few arguments...\n")
		return
	}

	// Take the first argument as path to CMakeLists.txt
	listFilePath := os.Args[1]

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	newTemplate := getTemplate(listFilePath)

	fmt.Print(newTemplate)
}

func getTemplate(listFilePath string) string {
	// Read files into strings
	contentList, err := ioutil.ReadFile(listFilePath)
	if err != nil {
		panic(err)
	}

	// Find library or executable names
	targets := findLibraryNames(string(contentList))
	for i, match := range targets {
		fmt.Printf("match %d: %s\n", i, match)
	}

	// Find project name
	projectName, err := findProjectName(string(contentList))
	fmt.Printf("Project name: %s\n", projectName)

	// Puts project name into template
	newTemplate := replaceString(template, `\$\{PROJECT_NAME\}`, projectName)

	// Puts project name into template
	return replaceString(newTemplate, `\$\{TARGETS\}`, strings.Join(targets, " "))
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

func replaceString(src string, pattern string, repl string) string {
	r := regexp.MustCompile(pattern)
	return r.ReplaceAllString(src, repl)
}

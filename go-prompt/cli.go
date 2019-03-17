package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"os"

	"io/ioutil"

	"github.com/c-bata/go-prompt"
)

func getSuggestionsPath(path string) []string {
	// remove the letters after and inclusive the last "/"
	var re = regexp.MustCompile(`\w*$`)
	path = re.ReplaceAllString(path, ``)

	info, err := os.Stat(path)
	if err != nil {
		return []string{}
	}

	var children []string
	if info.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return []string{}
		}
		// Strip "/" at the end of path
		if strings.HasSuffix(path, "/") {
			path = path[0 : len(path)-1]
		}

		for _, file := range files {
			children = append(children, path+"/"+file.Name())
		}
	}

	return children
}

func createCompleter(textList []string) prompt.Completer {

	completer := func(d prompt.Document) []prompt.Suggest {
		var s []prompt.Suggest
		for _, value := range textList {
			s = append(s, prompt.Suggest{Text: value, Description: "placeholder"})
		}

		children := getSuggestionsPath(d.GetWordBeforeCursor())

		for _, value := range children {
			s = append(s, prompt.Suggest{Text: value, Description: "path"})
		}

		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}

	return completer
}

func main() {

	var options []string
	//	fi, err := os.Stdin.Stat()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	if fi.Mode()&os.ModeNamedPipe == 0 {
	//		fmt.Println("no pipe :(")
	//		return
	//	}
	//
	//	scanner := bufio.NewScanner(os.Stdin)
	//	for scanner.Scan() {
	//		options = append(options, scanner.Text())
	//	}
	//
	//	if err := scanner.Err(); err != nil {
	//		log.Println(err)
	//	}

	fmt.Println("Please select table.")
	t := prompt.Input("> ", createCompleter(options))
	fmt.Println("You selected " + t)

	out, err := exec.Command("date").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out)

}

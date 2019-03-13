package main

import (
	"fmt"
	"os/exec"

	"bufio"
	"log"
	"os"

	"github.com/c-bata/go-prompt"
)

func createCompleter(textList []string) prompt.Completer {

	completer := func(d prompt.Document) []prompt.Suggest {
		var s []prompt.Suggest
		for _, value := range textList {
			s = append(s, prompt.Suggest{Text: value, Description: "placeholder"})
		}

		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}

	return completer
}

func main() {

	var options []string
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("no pipe :(")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		options = append(options, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println("Please select table.")
	t := prompt.Input("> ", createCompleter(options))
	fmt.Println("You selected " + t)

	out, err := exec.Command("date").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out)
}

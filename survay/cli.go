package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

func createQuestion(name string, message string, options []string) []*survey.Question {
	return []*survey.Question{
		{
			Name: name,
			Prompt: &survey.MultiSelect{
				Message: message,
				Options: options,
			},
		},
	}
}

func main() {
	// Get data options from pipe
	options := []string{
		"option 1",
		"option 2",
	}
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
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// ask the question
	answers := []string{}
	question := createQuestion("letter", "Select topics...", options)
	err = survey.Ask(question[:], &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("you chose: %s\n", strings.Join(answers, ", "))
}

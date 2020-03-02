package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	apiUrl := "https://httpbin.org/get"
	get_info(apiUrl)
}

func get_info(apiUrl string) string {
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		panic(err)
	}

	// For the example the credentials are not needed
	username, passwd := credentials()
	req.SetBasicAuth(username, passwd)

	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// TODO: replace ReadAll with something less dangerous
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", string(body))
	return string(body)
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		panic(err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

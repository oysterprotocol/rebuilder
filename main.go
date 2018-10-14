package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"time"
)

// Response structs
type checkStatusRes struct {
	Available bool `json:"available"`
}

// TODO:  update this when "status" is moved out of v2
const statusPath = "http://localhost:3000" +
	string(os.PathSeparator) + "api" +
	string(os.PathSeparator) + "v2" +
	string(os.PathSeparator) + "status"

// TODO:  make this more configurable rather than hardcoding it
const brokerEnvPath = string(os.PathSeparator) + "home" +
	string(os.PathSeparator) + "ubuntu" +
	string(os.PathSeparator) + "brokernode" +
	string(os.PathSeparator) + ".env"

func main() {
	fmt.Println("___________________________________________________")
	fmt.Println("Checking broker status on path: " + statusPath)
	statusResponse, err := http.Get(statusPath)
	fmt.Println("statusResponse: ")
	fmt.Println(statusResponse)

	if err != nil {
		fmt.Println("Status check error: ")
		fmt.Println(err)
		rebuild()
	} else {
		statusResParsed := &checkStatusRes{}
		parseErr := parseResBody(statusResponse, statusResParsed)
		if parseErr != nil {
			fmt.Println("Parsing error: ")
			fmt.Println(parseErr)
			rebuild()
		} else {
			fmt.Print("Parsed result, broker available: ")
			fmt.Println(statusResParsed.Available)
			fmt.Println("no need to rebuild")
		}
	}
	fmt.Println("___________________________________________________")
	defer statusResponse.Body.Close() // we need to close the connection
}

func rebuild() {
	currentTime := time.Now()

	input, err := ioutil.ReadFile(brokerEnvPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(input), "\n")

	lines = append(lines, "\n# Rebuild occurred: "+currentTime.Format("Mon Jan _2 15:04 2006"))

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(brokerEnvPath, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// parseResBody take a request and parses the body to the target interface.
func parseResBody(res *http.Response, dest interface{}) error {
	body := res.Body
	defer body.Close()

	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, dest)
	return err
}

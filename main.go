package main

import (
	"net"
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
const statusPath =
	":3000" +
	string(os.PathSeparator) + "api" +
	string(os.PathSeparator) + "v2"+
	string(os.PathSeparator) + "status"

// TODO:  make this more configurable rather than hardcoding it
const envPath =
	string(os.PathSeparator) + "home" +
	string(os.PathSeparator) + "ubuntu"+
	string(os.PathSeparator) + "brokernode" +
	string(os.PathSeparator) + ".env"

func main() {
	ipAddress := getLocalIP()

	if ipAddress != "" {
		statusResponse, err := http.Get(ipAddress + statusPath)
		defer statusResponse.Body.Close() // we need to close the connection
		if err != nil {
			fmt.Println(err)
			return
		}

		statusResParsed := &checkStatusRes{}
		if err := parseResBody(statusResponse, statusResParsed); err != nil {
			rebuild()
		}
	}
}

func rebuild() {
	currentTime := time.Now()

	input, err := ioutil.ReadFile(envPath)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	lines = append(lines, "\n# Rebuild occurred: " + currentTime.Format("Mon Jan _2 15:04 2006"))

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(envPath, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
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

// The UAVreg Registry Client provides an interface to the HTTP API of the built-in UAV registry
// of the UAVreg PKI Server.
//
// Copyright: 2021 Jannik Beyerstedt (consider it GmbH)

package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const buildVersion string = "v1.0.0"

// configuration parameters
const droneIDByteLen = 18

// constants and global variables
const droneIDHexLen = 2 * droneIDByteLen

var droneID string
var validityDurationWeeks int
var validityDurationHours int
var apiURL string

func init() {
	flag.StringVar(&droneID, "d", "0", "drone ID as hex string")
	flag.IntVar(&validityDurationWeeks, "w", 0, "validity duration in weeks, set 0 to delete")
	flag.IntVar(&validityDurationHours, "h", 1, "validity duration in hours, set 0 to delete")
	flag.StringVar(&apiURL, "u", "http://localhost:8080/registry", "URL to send the POST request to")
	flag.Parse()
}

func main() {
	fmt.Println("droneregistry client " + buildVersion)
	fmt.Println("- drone ID: ", droneID)
	fmt.Println("- validity: ", validityDurationWeeks, "weeks,", validityDurationHours, "hours")
	fmt.Println("- url:      ", apiURL)
	fmt.Println("")

	// check inputs
	if droneID == "0" {
		fmt.Println("Please set a drone id with the -d flag")
		log.Fatalln("fatal error: no droneID set")
	}
	if len(droneID) != droneIDHexLen {
		fmt.Println("Drone id must be " + strconv.Itoa(droneIDByteLen) + " Byte long")
		log.Fatalln("fatal error: wrong size of droneID")
	}

	// parse drone ID to byte array (to check if it is valid)
	droneIDBytes := make([]byte, droneIDByteLen)
	_, err := hex.Decode(droneIDBytes, []byte(droneID))
	if err != nil {
		fmt.Println("Drone id must contain valid hex chars")
		log.Fatalln("fatal error:", err)
	}

	// build API query string
	var validityHours = (7 * validityDurationWeeks) + validityDurationHours
	var jsonStr = []byte("")
	var httpMethod = ""

	if validityHours == 0 { // DELETE drone from registry
		httpMethod = "DELETE"
		jsonStr = []byte(`[{"droneId": "` + droneID + `"}]`)

	} else { // POST (update) drone validity
		httpMethod = "POST"

		// auto-set validity from now until now + $validityDurationWeeks weeks + $validityDurationHours weeks
		validFrom := time.Now()
		validUntil := time.Now().Add(time.Hour * time.Duration((validityDurationWeeks*7*24)+validityDurationHours))

		jsonStr = []byte(`[{"droneId": "` + droneID + `", "validFrom": "` + validFrom.Format(time.RFC3339) + `", "validUntil": "` + validUntil.Format(time.RFC3339) + `"}]`)
	}

	// send data to dronePKI
	fmt.Println("sending request to registry...")
	req, err := http.NewRequest(httpMethod, apiURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatalln("fatal error:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("fatal error:", err)
	}
	defer resp.Body.Close()

	// check for response status 200
	if strings.HasPrefix(resp.Status, "200") {
		fmt.Print("SUCCESS: ")
	} else {
		fmt.Println("ERROR:", resp.Status)
		fmt.Print(" >>>>  ")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

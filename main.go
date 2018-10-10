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

const buildVersion string = "v0.0.1"

// configuration parameters
const droneIDByteLen = 18

// constants and global variables
const droneIDHexLen = 2 * droneIDByteLen

var droneID string
var apiURL string

func init() {
	flag.StringVar(&droneID, "d", "0", "drone ID as hex string")
	flag.StringVar(&apiURL, "u", "http://localhost:8080/registry", "URL to send the POST request to")
	flag.Parse()
}

func main() {
	fmt.Println("droneregistry client " + buildVersion)
	fmt.Println("- drone ID: ", droneID)
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
		log.Fatalln("fatal error:\n", err)
	}

	// auto-set valdity from now until now + 2 months
	validFrom := time.Now()
	validUntil := time.Now().AddDate(0, 2, 0)

	var jsonStr = []byte(`[{"droneId": "` + droneID + `", "validFrom": "` + validFrom.Format(time.RFC3339) + `", "validUntil": "` + validUntil.Format(time.RFC3339) + `"}]`)

	// send data to dronePKI
	fmt.Println("sending request to registry...")
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("fatal error:\n", err)
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

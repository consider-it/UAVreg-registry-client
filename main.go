package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strconv"
)

const buildVersion string = "v0.0.1"

// configuration parameters
const droneIDByteLen = 18

// constants and global variables
const droneIDHexLen = 2 * droneIDByteLen

var droneID string
var revoke bool

func init() {
	flag.StringVar(&droneID, "d", "0", "drone ID as hex string")
	flag.BoolVar(&revoke, "r", false, "revoke insurance for this drone")
	flag.Parse()
}

func main() {
	fmt.Println("droneregistry client " + buildVersion)
	fmt.Println("- drone ID: ", droneID)
	fmt.Println("- revoke?:  ", revoke)
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

	// TODO: actually send data to dronePKI
	log.Fatalln("TODO: nothing implemented yet")
}

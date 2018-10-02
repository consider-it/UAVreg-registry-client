package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

const buildVersion string = "v0.0.1"

const droneIdLenByte int = 18

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
		fmt.Println("please set a drone id with the -d flag")
		log.Fatalln("no droneID set")
	}
	if len(droneID) != 2*droneIdLenByte {
		fmt.Println("drone id must be 18 Byte long")
		log.Fatalln("wrong size of droneID")
	}

	// parse drone ID to byte array (to check if it is valid)
	droneIDBytes := make([]byte, droneIdLenByte)
	_, err := hex.Decode(droneIDBytes, []byte(droneID))
	if err != nil {
		fmt.Println("drone id must contain valid hex chars")
		log.Fatal(err)
	}

	// TODO: actually send data to dronePKI
	log.Fatal("TODO: nothing implemented yet")
}

# UAV e-Registration: PKI Server
_A demo PKI for drone certification used to sign D2X (Drone to anything) messages._

The PKI consists of three "server"-side applications, which inter-connect, as well as a possible third-party drone registry service.
The cryptographic signature algorithm used is ECDSA with the NIST P-256 curve.

## Information
This application is the main part, which issues certificate to drones (which have sent a CSR).
The implementation can use different drone registies and it also has a very simple registry built in.


## Installation
Because the application is written in [go](https://golang.org), a working go directory and compiler must be present.

To compile the binary yourself, download (or get if ssh to this bitbucket repo is set up) the application and it's dependencies:
```
cd ~/go
go get bitbucket.org/nxp-d2x/d2xregistryclient
cd bitbucket.org/nxp-d2x/d2xregistryclient
go get
```

Run any application without installing:
```
go run *.go $FLAGS
```

Install the application to `~/go/bin/`:
```
go install
```
(Don't forget to set the `$GOBIN` environment variable to `~/go/bin` and to add that to your `$PATH`.)


## Usage
To experiment with the internal drone registry (sqlite database), a simple insurance client is provided.
This client can register or update a drone's insurance validity.

To keep things simple, the validity period will always start now.
The available command line flags are (default value printed bold):

- `-d` Drone ID as hex string, required
- `-w` Validity duration in weeks
- `-u` API URL: **`http://localhost:8080/registry`**


## Architecture
see https://bitbucket.org/nxp-drone/d2xpkiserver

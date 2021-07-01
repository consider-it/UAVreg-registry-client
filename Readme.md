# UAV e-Registration: Demo UAV Registry Client
_A client to register UAVs in the built-in demo UAV registry of the UAVreg-PKI-server._


## Installation and Usage

### Installation
Please make sure, that a [go](https://golang.org) toolchain is present and your go working directory is set up properly.

Now go to `$GOPATH/src/` and clone this repository:
```shell
git clone gitea@office.consider-ip.com:UDVeo/UAVreg-registry-client.git
```

Enter the repositories directory and fetch the dependencies:
```shell
go get
```

Finally you can either run without installing or install the application to your go binary directory:
```shell
## just compile and run
go run *.go <flags>

## install to $HOME/go/bin/ and run from there
go install
$HOME/go/bin/UAVreg-registry-client <flags>
```
Don't forget to add the go binary path to your `$PATH` environment variable, if you like to execute the application like every other application.


### Usage
This client applications uses the API of the built-in demo UAV registry provided by the UAVreg-PKI-server.
So the server has to run when using this application.

Every UAV registration has a validity period and the UAV is identified by it's serial number (Drone ID).
To simplify the interface, the validity period will always start now and a duration can be specified, which will be converted to an absolute timestamp automatically.

The available command line flags are (default value printed bold):

- `-d` Drone ID as hex string, required
- `-w` Validity duration in weeks
- `-u` API URL: **`http://localhost:8080/registry`**

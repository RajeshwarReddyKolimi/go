# Go installation

## Steps followed

- Dowloaded go using wget `https://go.dev/dl/go1.24.1.linux-amd64.tar.gz`
- Extracted it and installed using `sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz`
- Set path using `export PATH=$PATH:/usr/local/go/bin`
- Verified installation using `go version`
- Created a folder goinstallation.
- Ran `go mod init goinstallation` which created a `go.mod` file.
- Created a `main.go` file with main function which logs "Hello, Go!".
- Ran the file using `go run main.go`.
- Built the file using `go build` which created a binary file named `goinstallation`.
- Ran the built file using `./goinstallation`

## Output

```
$ wget https://go.dev/dl/go1.24.1.linux-amd64.tar.gz

$ sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz

$ export PATH=$PATH:/usr/local/go/bin

$ go version
go version go1.24.1 linux/amd64

$ go mod init go-installation
go: creating new go.mod: module go-installation
go: to add module requirements and sums:
        go mod tidy

$ go run main.go
Hello, Go!

$ go build
$ ./go-installation
Hello, Go!
```

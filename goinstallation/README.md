# Go installation

## Steps followed

- Dowloaded go using wget `https://go.dev/dl/go1.24.1.linux-amd64.tar.gz`
- Extracted it and installed using `sudo tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz`
- Set path using `export PATH=$PATH:/usr/local/go/bin`
- Verified installation using `go version`
- Created a folder goinstallation.
- Run `go mod init goinstallation` which created a `go.mod` file.
- Created a `main.go` file with main function which logs "Hello, Go!".
- Run the file using `go run main.go`.
- Built the file using `go build` which created a binary file named `goinstallation`.
- Run the built file using `./goinstallation`

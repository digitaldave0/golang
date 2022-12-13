# golang
golang learnings

## Install Golang 

```bash
VERSION="1.19.4" # go version
ARCH="amd64" # go archicture
curl -O -L "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
sha256sum go1.16.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
sudo nano ~/.profile
export PATH=$PATH:/usr/local/go/bin
go version
```


## Test Install 
go mod init test/hello


```bash 
mkdir test
echo -e "package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}" > hello.go
go run .
```

```bash
Output
Hello, World!
```

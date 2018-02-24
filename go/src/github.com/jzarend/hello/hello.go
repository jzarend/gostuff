package main

import (
        "fmt"
        "github.com/jzarend/stringutil"
)

func main() {
    fmt.Printf("Hello World!\n")
    fmt.Printf(stringutil.Reverse("Hello World!")+"\n")
    fmt.Printf(stringutil.Reverse("!dlroW olleH")+"\n")
}

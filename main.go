package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := strings.Fields(os.Args[1])
	fmt.Println(len(a)+1)
}

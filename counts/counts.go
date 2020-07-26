package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var n int

	for i := 1; i < len(os.Args); i++ {
		fn := os.Args[i]
		text, err := ioutil.ReadFile(fn)

		if err != nil {
			fmt.Fprintf(os.Stderr, "can't read %s: %s\n", fn, err)
			continue
		}

		words := strings.Fields(string(text))
		n += len(words)
	}

	fmt.Println(n, "total words")
}

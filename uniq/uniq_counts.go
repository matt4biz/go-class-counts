package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int)

	for i := 1; i < len(os.Args); i++ {
		fn := os.Args[i]
		text, err := ioutil.ReadFile(fn)

		if err != nil {
			fmt.Fprintf(os.Stderr, "can't read %s: %s\n", fn, err)
			continue
		}

		words := strings.Fields(string(text))

		for _, w := range words {
			m[w] += 1
		}
	}

	fmt.Println(len(m), "unique words")
}

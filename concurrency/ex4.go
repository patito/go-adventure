package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "o nani peida na mao e cheira com a boca"
	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Println(<-ch + " ")

	}

}

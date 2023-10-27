package main

import (
	paf "github.com/44/paf/src"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(paf.FormatGrep(text))
		// fmt.Println(text)
	}
}


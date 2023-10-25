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
		fmt.Println(paf.FormatGrep(scanner.Text()))
	}
}


package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		fmt.Println(os.Args)
		for i, arg := range os.Args {
			fmt.Println("%s at %s", arg, i)
		}
	}
}

func readCsvToString() {
	//first check for a size shorter than a limit
}

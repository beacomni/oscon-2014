package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("sorry fellers, need a file name")
		return
	}
	filename := os.Args[1]
	valuesToSort := extractSliceFromCsv(filename)
	MergeSort(valuesToSort)
}

func MergeSort(valuesToSort){
	//try with making a goroutine, and without
	firstHalf := valuesToSort[0]

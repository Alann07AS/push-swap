package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"pushSwap/funcSw"
)

func main() {
	stackA, stackB := funcSw.InitFuncSw()
	args := os.Args
	if len(args) == 1 {
		log.Fatal("No args")
	}
	// fild stackA
	for _, each := range args[1:] {
		in, err := strconv.Atoi(each)
		if err != nil {
			log.Fatal(err)
		}
		*stackA = append(*stackA, in)
	}

	// fild StackB
	*stackB = make([]int, 0)
	// check if is dublicate
	isDuplicate(*stackA)

	// program SELECTION SORT
	isMod := true
	for isMod {
		funcSw.PB()
		for range *stackA {
			if funcSw.GetItemA(0) >= funcSw.GetItemB(0) {
				funcSw.PB()
			} else {
				funcSw.RA()
			}
		}
	}

	funcSw.PrintLog()
	fmt.Println(stackA)
}

func isDuplicate(arr []int) {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num]++
		if dict[num] > 1 {
			log.Fatal("Error Repetition")
		}
	}
}

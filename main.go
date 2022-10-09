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
	*stackB = make([]int, len(*stackA))
	funcSw.RRA()
	fmt.Println(*stackA, *stackB)
}

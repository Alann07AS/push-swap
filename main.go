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
	isDuplicate(*stackA)
	// test:   go run . 2 1 3 6 8 5
	funcSw.SA()
	funcSw.PB()
	funcSw.PB()
	funcSw.PB()
	funcSw.RB()
	funcSw.RRR()
	funcSw.PA()
	funcSw.PA()
	funcSw.PA()
	fmt.Println(*stackA, *stackB)
	// result : [1 2 3 5 6 8] []
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

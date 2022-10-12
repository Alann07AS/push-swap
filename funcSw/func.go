package funcSw

import (
	"fmt"
	"strings"
)

// all Push Swap func_________________

var stackA, stackB []int

// init toutes les functions
func InitFuncSw() (*[]int, *[]int) {
	return &stackA, &stackB
}

// push index 0 de stackA sur stackB
func PA() {
	pa.logFunc()
	push(&stackA, &stackB)
}

// push index 0 de stackB sur stackA
func PB() {
	pb.logFunc()
	push(&stackB, &stackA)
}

func push(st1, st2 *[]int) {
	stack1, stack2 := *st1, *st2
	if len(stack2) > 0 {
		*st1 = append([]int{stack2[0]}, stack1...)
		*st2 = stack2[1:]
	}
}

// swap the two first element of stackA
func SA() {
	sa.logFunc()
	swap(stackA)
}

// swap the two first element of stackB
func SB() {
	sb.logFunc()
	swap(stackB)
}

func swap(stack []int) {
	if len(stack) > 1 {
		stack[0], stack[1] = stack[1], stack[0]
	}
}

// exec SA et SB
func SS() {
	ss.logFunc()
	SA()
	SB()
	removeTwoLastLog()
}

// shift up elements du stack A, le premier devient le derier
func RA() {
	ra.logFunc()
	shiftUp(&stackA)
}

// shift up elements du stack B, le premier devient le derier
func RB() {
	rb.logFunc()
	shiftUp(&stackB)
}

func shiftUp(stack *[]int) {
	st := *stack
	*stack = append(st[1:], st[0])
}

// exec RA et RB
func RR() {
	rr.logFunc()
	RA()
	RB()
	removeTwoLastLog()
}

// shift down elements du stack A, le dernier devient le premier
func RRA() {
	rra.logFunc()
	shiftDown(&stackA)
}

// shift down elements du stack B, le dernier devient le premier
func RRB() {
	rrb.logFunc()
	shiftDown(&stackB)
}

func shiftDown(stack *[]int) {
	st := *stack
	l := len(st) - 1
	*stack = append([]int{st[l]}, st[:l]...)
}

// exec RRA et RRB
func RRR() {
	rrr.logFunc()
	RRA()
	RRB()
	removeTwoLastLog()
}

// get Index A
func GetItemA(i int) int {
	return getItem(stackA, i)
}

// get Index B
func GetItemB(i int) int {
	return getItem(stackB, i)
}

func getItem(stack []int, i int) int {
	if len(stack) != 0 {
		return stack[i]
	}
	return 0
}

//______________________________//

// log compt func call___________

type SwConst string

const (
	pa  SwConst = "pa\n"
	pb  SwConst = "pb\n"
	sa  SwConst = "sa\n"
	sb  SwConst = "sb\n"
	ss  SwConst = "ss\n"
	ra  SwConst = "ra\n"
	rb  SwConst = "rb\n"
	rr  SwConst = "rr\n"
	rra SwConst = "rra\n"
	rrr SwConst = "rrr\n"
	rrb SwConst = "rrb\n"
)

var funcLog = []string{}

func (f SwConst) logFunc() {
	funcLog = append(funcLog, string(f))
}

func removeTwoLastLog() {
	funcLog = funcLog[:len(funcLog)-2]
}

// Print all func used
func PrintLog() {
	fmt.Print(strings.Join(funcLog, ""))
}
//________________________//

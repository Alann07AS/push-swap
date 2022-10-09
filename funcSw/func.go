package funcSw

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

// swap the two first elemnt of stackA
func SA() {
	sa.logFunc()
	if len(stackA) > 1 {
		swap(stackA)
	}
}

// swap the two first elemnt of stackB
func SB() {
	sb.logFunc()
	if len(stackB) > 1 {
		swap(stackB)
	}
}

func swap(stack []int) {
	stack[0], stack[1] = stack[1], stack[0]
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

//______________________________//

// log compt func call___________

type SwConst string

const (
	pa  SwConst = "pa"
	pb  SwConst = "pb"
	sa  SwConst = "sa"
	sb  SwConst = "sb"
	ss  SwConst = "ss"
	ra  SwConst = "ra"
	rb  SwConst = "rb"
	rr  SwConst = "rr"
	rra SwConst = "rra"
	rrr SwConst = "rrr"
	rrb SwConst = "rrb"
)

var funcLog = []string{}

func (f SwConst) logFunc() {
	funcLog = append(funcLog, string(f))
}

func removeTwoLastLog() {
	funcLog = funcLog[:len(funcLog)-2]
}

// get Log var pointer
func GetLog() *[]string {
	return &funcLog
}

//________________________//

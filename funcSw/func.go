package funcSw

var stackA, stackB []int

// init toutes les functions
func InitFuncSw() (*[]int, *[]int) {
	return &stackA, &stackB
}

// push index 0 de stackA sur stackB
func PA() {
	stackA[0] = stackB[0]
}

// push index 0 de stackB sur stackA
func PB() {
	stackB[0] = stackA[0]
}

// swap the two first elemnt of stackA
func SA() {
	swap(stackA)
}

// swap the two first elemnt of stackB
func SB() {
	swap(stackB)
}

func swap(stack []int) {
	stack[0], stack[1] = stack[1], stack[0]
}

// exec SA et SB
func SS() {
	SA()
	SB()
}

// shift up elements du stack A, le premier devient le derier
func RA() {
	shiftUp(&stackA)
}

// shift up elements du stack B, le premier devient le derier
func RB() {
	shiftUp(&stackB)
}

func shiftUp(stack *[]int) {
	st := *stack
	*stack = append(st[1:], st[0])
}

// exec RA et RB
func RR() {
	RA()
	RB()
}

// shift down elements du stack A, le dernier devient le premier
func RRA() {
	shiftDown(&stackA)
}

// shift down elements du stack B, le dernier devient le premier
func RRB() {
	shiftDown(&stackB)
}

func shiftDown(stack *[]int) {
	st := *stack
	l := len(st) - 1
	*stack = append([]int{st[l]}, st[:l]...)
}

// exec RRA et RRB
func RRR() {
	RRA()
	RRB()
}

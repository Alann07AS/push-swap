package funcSw

var stackA, stackB []int

// init toutes les functions
func InitFuncSw() (*[]int, *[]int) {
	return &stackA, &stackB
}

// push index 0 de stackA sur stackB
func PA() {
	push(&stackA, &stackB)
}

// push index 0 de stackB sur stackA
func PB() {
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
	if len(stackA) > 1 {
		swap(stackA)
	}
}

// swap the two first elemnt of stackB
func SB() {
	if len(stackB) > 1 {
		swap(stackB)
	}
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

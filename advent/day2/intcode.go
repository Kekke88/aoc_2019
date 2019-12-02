package day2

type IntcodeComputer struct {
	Memory          []int
	OperationOffset int
}

const END_OF_PROGRAM = 99
const FUNC_ADDITION = 1
const FUNC_MULTIPLY = 2
const SIZE_OF_INSTRUCTION_SET = 4

const MEM_POS_SUM = 0
const MEM_POS_ONE_OFFSET = 1
const MEM_POS_TWO_OFFSET = 2
const MEM_POS_SUM_OFFSET = 3

func create(mem []int) *IntcodeComputer {
	// Init new IntcodeComputer with array and set CO to 0
	c := IntcodeComputer{Memory: mem, OperationOffset: 0}

	return &c
}

func (c IntcodeComputer) Run() {
	// Run through commands
	for {
		if c.CurrentOperation() == END_OF_PROGRAM {
			break
		}

		switch c.CurrentOperation() {
		case FUNC_ADDITION:
			c.SetMemState(c.GetSumPos(), c.GetValue(1)+c.GetValue(2))
		case FUNC_MULTIPLY:
			c.SetMemState(c.GetSumPos(), c.GetValue(1)*c.GetValue(2))
		default:
			panic("Operation not found")
		}

		c.NextOp()
	}
}

func (c IntcodeComputer) CurrentOperation() int {
	return c.Memory[c.OperationOffset]
}

func (c IntcodeComputer) SetMemState(p int, v int) {
	c.Memory[p] = v
}

func (c IntcodeComputer) GetValue(pos int) int {
	return c.Memory[c.Memory[c.OperationOffset+pos]]
}

func (c IntcodeComputer) GetSumPos() int {
	return c.Memory[c.OperationOffset+MEM_POS_SUM_OFFSET]
}

func (c *IntcodeComputer) NextOp() {
	c.OperationOffset += SIZE_OF_INSTRUCTION_SET
}

func (c IntcodeComputer) Value() int {
	// Returns value
	return c.Memory[MEM_POS_SUM]
}

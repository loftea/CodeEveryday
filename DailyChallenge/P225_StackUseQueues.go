type MyStack struct {
	q_temp []int
	q_main []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q_temp = append(this.q_temp, x)
	for len(this.q_main) > 0 {
		this.q_temp = append(this.q_temp, this.q_main[0])
		this.q_main = this.q_main[1:]
	}

	for len(this.q_temp) > 0 {
		this.q_main = append(this.q_main, this.q_temp[0])
		this.q_temp = this.q_temp[1:]
	}

}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.q_main[0]
	this.q_main = this.q_main[1:]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q_main[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q_main) == 0
}
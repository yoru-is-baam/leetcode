package main

import "fmt"

type MinStack struct {
	minStack []int
	stack    []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	if len(this.minStack) == 0 || value <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, value)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}

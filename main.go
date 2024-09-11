package main

import "fmt"

type Stack struct {
	elements []int
}

func main() {

	stack := Stack{}

	stack.push(2)
	stack.push(5)
	stack.push(9)
	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.push(40)
	stack.push(50)

	fmt.Println("Pilha after push`s")
	for _, e := range stack.elements { //Complexidade de tempo Big O(n)
		fmt.Println(e)
	}

	fmt.Println("Peek: ", stack.peek()) // Complexidade de tempo Big O(1) constante

	stack.pop()
	stack.pop()
	stack.pop()

	fmt.Println("Pilha after pop`s")
	for _, e := range stack.elements { //Complexidade de tempo Big O(n)
		fmt.Println(e)
	}

	fmt.Println("Peek: ", stack.peek())

	fmt.Println("Sum: ", stack.sum())
	fmt.Println("LargestNum: ", stack.largestNum())
	fmt.Println("SmallestNum: ", stack.smallestNum())
}

func (m *Stack) pop() (element int) {

	if len(m.elements) == 0 { //Complexidade de tempo Big O(1) constante
		return 0
	}

	lastElement := len(m.elements) - 1    //Complexidade de tempo Big O(1) constante
	element = m.elements[lastElement]     //Complexidade de tempo Big O(1) constante
	m.elements = m.elements[:lastElement] //Complexidade de tempo Big O(1) constante

	return element
}

func (m *Stack) push(element int) {
	m.elements = append(m.elements, element) //Complexidade de tempo Big O(1) constante
}

func (m *Stack) largestNum() (num int) {
	if len(m.elements) == 0 { //Complexidade de tempo Big O(1) constante
		return 0
	}

	num = m.elements[0] //Complexidade de tempo Big O(1) constante

	for _, e := range m.elements { //Complexidade de tempo Big O(n)
		if e > num {
			num = e
		}
	}

	return num
}

func (m *Stack) smallestNum() (result int) {
	if len(m.elements) == 0 { //Complexidade de tempo Big O(1) constante
		return 0
	}

	result = m.elements[0] //Complexidade de tempo Big O(1) constante

	for _, e := range m.elements { //Complexidade de tempo Big O(n)
		if e < result {
			result = e
		}
	}

	return result
}

func (m *Stack) sum() (result int) {
	if len(m.elements) == 0 { //Complexidade de tempo Big O(1) constante
		return 0
	}

	for _, e := range m.elements { //Complexidade de temp Big O(n) Tempo Linear
		result += e
	}

	return result
}

func (m *Stack) peek() (retVal int) {
	if len(m.elements) == 0 {
		return 0
	}

	return m.elements[len(m.elements)-1]
}

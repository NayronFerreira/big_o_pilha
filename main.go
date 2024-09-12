package main

import "fmt"

type Stack struct {
	elements []int
}

func main() {

	stack := Stack{}

	stack.push(1)
	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.push(40)
	stack.push(50)
	stack.push(60)
	stack.push(70)
	stack.push(80)
	stack.push(90)

	fmt.Println("Stack after pushs")
	for _, e := range stack.elements {
		fmt.Println(e)
	}

	fmt.Println("Pop: ", stack.pop())
	fmt.Println("Pop: ", stack.pop())
	fmt.Println("Pop: ", stack.pop())
	fmt.Println("Pop: ", stack.pop())

	fmt.Println("Stack after pops")
	for _, e := range stack.elements {
		fmt.Println(e)
	}

	fmt.Println("Peek: ", stack.peek())
	fmt.Println("largest: ", stack.largestNum())
	fmt.Println("smallest: ", stack.smallestNum())
	fmt.Println("sum: ", stack.sum())

}

// verifica se a pilha esta vazia
func (m *Stack) isEmpty() bool {
	return len(m.elements) == 0 //Complexidade espaco tempo Big O(1)
}

// Retorna e remove o topo
func (m *Stack) pop() (retVal int) {

	if m.isEmpty() { //Complexidade espaco tempo Big O(1) constante
		return 0
	}

	lastElement := len(m.elements) - 1    //Complexidade espaco tempo Big O(1) constante
	retVal = m.elements[lastElement]      //Complexidade espaco tempo Big O(1) constante
	m.elements = m.elements[:lastElement] //Complexidade espaco tempo Big O(1) constante

	return retVal
}

// Retorna o topo
func (m *Stack) peek() (retVal int) {
	if m.isEmpty() { //Complexidade espaco tempo Big O(1) constante
		return 0
	}

	lastElement := len(m.elements) - 1 //Complexidade espaco tempo Big O(1) constante
	retVal = m.elements[lastElement]   //Complexidade espaco tempo Big O(1) constante

	return retVal
}

// adiciona elemento ao topo
func (m *Stack) push(element int) {
	m.elements = append(m.elements, element) //Complexidade espaco tempo Big O(1) constante
}

// retorna o tamanho da pilha
func (m *Stack) size() (retVal int) {
	return len(m.elements) //Complexidade espaco tempo Big O(1) constante
}

// Retorna o maior valor da Lista
func (m *Stack) largestNum() (retVal int) {
	if m.isEmpty() { //Complexidade espaco tempo Big O(1) constante
		return 0
	}

	retVal = m.elements[0] //Complexidade espaco tempo Big O(1) constante

	for _, e := range m.elements { //Complexidade espaco tempo Big O(n) linear
		if e > retVal {
			retVal = e
		}
	}

	return retVal
}

// Retorna o menor valor da Lista
func (m *Stack) smallestNum() (retVal int) {
	if m.isEmpty() { //Complexidade espaco tempo Big O(1) constante
		return 0
	}

	retVal = m.elements[0] //Complexidade espaco tempo Big O(1) constante

	for _, e := range m.elements { //Complexidade espaco tempo Big O(n) linear
		if e < retVal {
			retVal = e
		}
	}

	return retVal
}

// retorna a soma dos itens da pilha
func (m *Stack) sum() (retVal int) {
	if m.isEmpty() { //Complexidade espaco tempo Big O(1) constante
		return 0
	}

	for _, e := range m.elements { //Complexidade espaco tempo Big O(n) linear
		retVal += e
	}

	return retVal
}

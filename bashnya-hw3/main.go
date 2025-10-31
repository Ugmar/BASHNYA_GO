package main

import(
	"fmt"
)

const (
	MIN_SIZE_STACK = 100
)

type Stack struct{
	arr []int
}

func (stack *Stack) Push(value int){
	stack.arr = append(stack.arr, value)
}

func (stack *Stack) Pop() (value int, err bool){
	if stack.Size() == 0{
		return 0, true
	}

	value = stack.arr[stack.Size() - 1]
	stack.arr = stack.arr[:stack.Size() - 1]

	return value, false
}

func (stack *Stack) Size() int{
	return len(stack.arr)
}

func (stack *Stack) IsEmpty() bool{
	return len(stack.arr) == 0
}

func (stack *Stack) Init(){
	stack.arr = nil
	stack.arr = make([]int, 0, MIN_SIZE_STACK)
}

func (stack *Stack) Clear(){
	stack.arr = stack.arr[:0]
}

func main(){
	var stack Stack

	fmt.Printf("Стак не инициализировн\n")
	fmt.Printf("Стек пуст: %t\n", stack.IsEmpty())

	(&stack).Init()
	fmt.Println("Стек инициализирован")

	for i := range 10{
		fmt.Printf("Элемент добавлен в стек: %d\n", i)
		(&stack).Push(i)
	}

	fmt.Printf("Стек пут: %t Длина стека: %d\n", stack.IsEmpty(), stack.Size())

	value, err := (&stack).Pop()
	fmt.Printf("Элемент успешно излвечен: %t Извлеченный элемент: %d\n", err, value)
	fmt.Printf("Длина стека: %d\n", stack.Size())

	(&stack).Clear()
	fmt.Println("Стек очищен")
	fmt.Printf("Длина стека: %d \n", stack.Size())
}

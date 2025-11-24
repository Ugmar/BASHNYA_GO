package main

import (
	"fmt"
)

func Conveyor(numbers []int){
	input := make(chan int)
	output := make(chan int)
	
	go func(){
		defer close(input)

		for _, x := range numbers{
			input <- x
		}
	}()

	go func(){
		defer close(output)

		for x := range input{
			output <- x * 2
		}
	}()

	for x := range output{
		fmt.Println(x)
	}
}

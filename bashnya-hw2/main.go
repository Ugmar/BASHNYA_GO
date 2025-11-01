package main

import "fmt"

const MAIN_DIGIT = 12307

func main(){
	var digit int
	fmt.Print("Введите чило: ")
	n, err := fmt.Scan(&digit)

	if n != 1 || err != nil{
		fmt.Printf("Ошибка ввода числа!!!")
		return
	}

	if digit >= MAIN_DIGIT{
		fmt.Printf("Введеное число больше %d\n", MAIN_DIGIT)
		fmt.Printf("Обрабатываемое число: %d %% %d + 1 = %d\n", digit, MAIN_DIGIT, digit % MAIN_DIGIT + 1)
		digit = digit % MAIN_DIGIT + 1
	}

	for digit < MAIN_DIGIT{
		if digit < 0{
			digit *= -1
		} else if digit % 7 == 0{
			digit *= 39
		} else if digit % 9 == 0{
			digit *= 19
			digit++
			continue
		} else{
			digit += 2
			digit *= 3
		}

		if digit % (13 * 9) == 0{
			fmt.Printf("service error")
			return
		} else {
			digit++
		}
	}

	fmt.Printf("Ваше исходное супер пупер число после всех преобразований: %d\n", digit)
}

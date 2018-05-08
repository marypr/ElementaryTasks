// task_7
package main

import (
	"fmt"
	"strconv"
)

func main() {
	

	fmt.Println(Fib(6, 120))

}
func Fib (numbers ...int) (result []int) { //swıtch + 2 functıons or map
	if len(numbers) == 2 {

	min:= numbers[0]
	max:=  numbers[1]

			result := make([]int, 0)
			a := 1
			b := 1
			x := 0

			for x < max{
				if x<max{
			x = a + b
			a = b
			b = x
			if x > min{
			result = append(result, x)
		}}
		}

			return result
	} else if len(numbers) ==1{

		result := make([]int, 0)
		a := 1
		b := 1
		x := 0


		for  {
			x = a + b
			a = b
			b = x
			s := strconv.Itoa(x)
			if len(s)==numbers[0]{
			result = append(result, x)}  else if len(s)==numbers[0]+1{
					break
		}
			}

		return result
	}
	return result
	}



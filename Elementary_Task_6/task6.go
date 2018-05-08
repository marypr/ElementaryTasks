// lab_6
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() { //
	var n, m string
	fmt.Println("Enter length")
	fmt.Scan(&n)
	length := checkers(n)
	fmt.Println("Enter min value")
	fmt.Scan(&m)
	minvalue := checkers(m)
	fmt.Println(NumSequence(length, minvalue))
}

func checkers(a string) int64 {
	q, err := strconv.ParseInt(a, 10, 64)
	if err != nil {

		fmt.Print("status: failed, reason: ", err)
	}
	if err == nil {
		if q < 0 {
			q = q * (-1)
		}
		if q == 0 {
			fmt.Println("status: failed, reason: number must be > 0")
			os.Exit(0)//убрать
		}
	}
	return q

}

func NumSequence(length, minvalue int64) string {
	var sequence string
	var i int64
	start := int64(math.Ceil(math.Sqrt(float64(minvalue))))
	for i = 0; i < length; i++ {
		sequence += strconv.Itoa((int(start + i))) //исправить
		sequence += " "

	}

	return sequence
}

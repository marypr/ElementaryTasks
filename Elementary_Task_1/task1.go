// Elementary_Tasks_1 project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sym string
	var leng, wids string

	fmt.Println("Please, enter the length of the desk")
	fmt.Scan(&leng)
	len:= checkers(leng)
	fmt.Println("Please, enter the width of desk")
	fmt.Scan(&wids)
	wid:= checkers(wids)
	fmt.Println("Please, enter the symbol")
	fmt.Scan(&sym)
	fmt.Println("The desk:")
	fmt.Println(draw_desk(len, wid, sym))

}

func checkers(a string) int64 {
	q, err := strconv.ParseInt(a,10, 64)
	if err != nil {
		fmt.Print("status: failed, reason: ", err)
	}
	if err == nil {
		if q<0{
			q = q*(-1)
		}

	}
	return q

}

func draw_desk(length int64, width int64, symbol string) string {
	var desk string
	var i, j int64
	for i = 0; i < length; i++ {
		for j = 0; j < width; j++ {
			switch {
			case (i+j)%2 == 0:
				desk += (symbol)
			default:
				desk += (" ")
			}
		}
		desk += "\n"
	}
	return desk
}

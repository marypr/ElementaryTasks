// Elementary_Task_4 project main.go
package main

import (
	"fmt"
)

func main() { //форматирование
	var stringToExplore string
	i:=0
	fmt.Println("Please, enter any string")
	fmt.Scan(&stringToExplore)
	length := len(stringToExplore) //rune

	for offset := 0; offset < length; offset++ {

		for substringLength := 1; substringLength <= length-offset; substringLength++ {

			substring := stringToExplore[offset : offset+substringLength]

			if isPalindrom(substring) && len(substring) > 1  { //i<1 - для вывода только первого палиндрома
i++
			fmt.Println("result", substring)

				}
		}
	}
	//fmt.Println("No more palindromes were found")

}
func isPalindrom(stringToTest string) bool {
	length := len(stringToTest)
	halfLength := length / 2

	for i := 0; i <= halfLength; i++ {
		if stringToTest[i] != stringToTest[length-i-1] {

			return false
		}
	}
	return true

}

// task5
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var start, end string
	fmt.Println("Please, enter starting number")
	fmt.Scan(&start)
	starting_number := checkers(start)
	fmt.Println("Please, enter ending number")
	fmt.Scan(&end)
	ending_number := checkers(end)

	if !Validate(starting_number, ending_number) {
		fmt.Println(Validate(starting_number, ending_number))
	} else {
		BestCountingSuccessTickets(starting_number, ending_number)
	}
}

func Validate(min, max uint64) bool {//разделить
	var checker bool

	if min > max {
		checker = false
		fmt.Println("Min-", min, "must be less than max - ", max)
		main()
	}
	if min > 999999 || min <= 0 || min > 999999 {
		checker = false
		fmt.Println("Min", min, "and Max", max,
			"must be in range from 1 to 999999")
	} else {
		checker = true

	}
	return checker
}

func checkers(a string) uint64 {
	if len(a) != 6 {
		fmt.Println("status: failed, reason: ticket must have 6 numbers")
		main()
	}
	q, err := strconv.ParseUint(a, 10, 64)
	if err != nil {

		fmt.Print("status: failed, reason: ", err)
	}

	return q

}

func BestCountingSuccessTickets(min, max uint64) {
	easyMethodCounter := 0
	hardMethodCounter := 0
	for i := min; i <= max; i++ {
		ticket := partition(i)

		if easyMethod(ticket) {
			easyMethodCounter++
		}

		if hardMethod(ticket) {
			hardMethodCounter++
		}
	}
	if easyMethodCounter > hardMethodCounter {
		fmt.Printf("Easy method won! Was found %d lucky tickets\r\n",
			easyMethodCounter)
		fmt.Println("Hard-", hardMethodCounter, "< Easy-", easyMethodCounter)
		os.Exit(0)
	} else if easyMethodCounter < hardMethodCounter {
		fmt.Printf("Hard method won! Was found %d lucky tickets\r\n",
			hardMethodCounter)
		fmt.Println("Hard-", hardMethodCounter, "> Easy-", easyMethodCounter)
		os.Exit(0)
	} else if easyMethodCounter == hardMethodCounter {
		fmt.Printf("They are both equal, 1 - %d lucky tickets, 2 - %d tickets\r\n", easyMethodCounter, hardMethodCounter)
		os.Exit(0)
	}

}

func partition(num uint64) (ticket [6]uint64) {
	for i := 5; i >= 0; i-- {
		ticket[i] = uint64(num % 10)
		num /= 10
	}
	return
}

func easyMethod(num [6]uint64) bool {
	return num[0]+num[1]+num[2] == num[3]+num[4]+num[5]
}
func hardMethod(num [6]uint64) bool {
	OddSum, EvenSum := 0, 0 //убрать большие буквы
	for _, i := range num {
		if (i % 2) == 0 {
			EvenSum += int(i)
		} else {
			OddSum += int(i)
		}
	}
	return OddSum == EvenSum
}

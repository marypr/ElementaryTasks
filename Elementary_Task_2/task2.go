// Elementary_Task_2 project main.go
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"os"
)



func main() {
	var w1, h1, w2, h2 float64
	rand.Seed(time.Now().UnixNano())
	var widthOfEnvelope, heightOfEnvelope string
	for { // исправить
		fmt.Println("Please, enter the width of the 1st envelope")
		fmt.Scan(&widthOfEnvelope)
		w1=checkers(widthOfEnvelope)
			//break
		}


	for {
		fmt.Println("Please, enter the height of the 1st envelope")
		fmt.Scan(&heightOfEnvelope)
		h1=checkers(heightOfEnvelope)
			break
		}

	for {
		fmt.Println("Please, enter the width of the 2st envelope")
		fmt.Scan(&widthOfEnvelope)
		w2=checkers(widthOfEnvelope)
			break
		}

	for {
		fmt.Println("Please, enter the height of the 2st envelope")
		fmt.Scan(&heightOfEnvelope)
		h2=checkers(heightOfEnvelope)
			break
		}

	env1 := envelope{width: w1, height: h1}
	env2 := envelope{width: w2, height: h2}
	//fmt.Println(calculations(env1, env2))
	fmt.Println(env1.calculations(env1,env2))

}

func checkers(a string) float64 {
	q, err := strconv.ParseFloat(a,64)
	if err != nil {
		fmt.Print("status: failed, reason: ", err)
		os.Exit(0)
	}
	if err == nil {
		if q < 0 {
			fmt.Println("u have entered negative number, but it's okay, i'll fix it")
			q = q * (-1)
		} else if q == 0 {
			fmt.Println("u have entered zero, so i'll replace this zero with any number i want")
			q = float64(rand.Intn(210))

	}
	}
	return q

}

type envelope struct {
	width  float64
	height float64
}

func (e envelope) calculations  (envel1 envelope, envel2 envelope) int { //исправить
	var n int
	area1 := envel1.height * envel1.width
	area2 := envel2.height * envel2.width

	switch {

	case area1 < area2:
		if (envel1.height < envel2.height && envel1.width < envel2.width) || (envel1.height < envel2.width && envel1.width < envel2.height) {
			n = 1
		}
	case area2 < area1:
		if (envel1.height > envel2.height && envel1.width > envel2.width) || (envel1.height > envel2.width && envel1.width > envel2.height) {
			n = 2
		}
	default:
		n = 0
	}
	return n
}

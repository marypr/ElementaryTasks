// Elementary_Task_3 project main.go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var names []Triangle
	fmt.Println("Please, enter parameters for triangles(split it with ',')",
		"to end - type 'next'")
	fmt.Println("Example - ABC, 10.5, 15, 12.4")

	input := bufio.NewScanner(os.Stdin)
	var re = regexp.MustCompile(`[[:space:]]`)
	for input.Scan() {

		if input.Text() == "next" {
			break
		}

		str := re.ReplaceAllString(input.Text(), "")
		s := strings.Split(str, ",")
		vertices, a, b, c := s[0], s[1], s[2], s[3]
		side1 := checkers(a)
		side2 := checkers(b)
		side3 := checkers(c)
		tring := Triangle{vertices: vertices, side1: side1, side2: side2, side3: side3}
		if Validate(tring) {

			names = append(names, tring)
		} else {
			fmt.Println("Such triangle just can't exist :(")
			fmt.Println("enter new:")
		}
	}
	//	names := []Triangle{
	//		{"ABC1", 12, 10, 8, 0},
	//		{"ABC2", 16, 10, 8, 0},
	//		{"ABC3", 13.64, 18, 8, 0},
	//	}

	Geron_calc(names)

	sort.Slice(names, func(i, j int) bool {
		return names[i].area > names[j].area
	})
	for index := range names {
		fmt.Println(names[index].vertices, names[index].area)
	}

}
func checkers(a string) float64 {
	q, err := strconv.ParseFloat(a, 64)
	if err != nil {

		fmt.Println("status: failed, reason: ", err)
	}
	if err == nil {

	}
	return q

}

type Triangle struct {
	vertices            string
	side1, side2, side3 float64
	area                float64
}

func Geron_calc(triangles []Triangle) {
	var area float64
	for index := range triangles {

		halfperimeter := ((triangles[index].side1 + triangles[index].side2 + triangles[index].side3) / 2)
		area = math.Sqrt(halfperimeter * ((halfperimeter - triangles[index].side1) * (halfperimeter - triangles[index].side2) * (halfperimeter - triangles[index].side3)))
		triangles[index].area = area
	}

}

func Validate(t Triangle) bool {
	isPositive := t.side1 > 0 && t.side2 > 0 && t.side3 > 0
	isCorrectSides := (t.side1 < t.side2+t.side3) &&
		(t.side2 < t.side1+t.side3) && (t.side3 < t.side1+t.side2)
	return isPositive && isCorrectSides
}

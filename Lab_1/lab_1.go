package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var A float64
	var B float64
	var C float64

	args := os.Args[1:]

	if len(args) > 0 {
		a, err := strconv.ParseFloat(args[0], 64)
		if err == nil {
			A = a
		} else {
			//fmt.Print(a)
			fmt.Println("Чёт A плохой, давай заново")
			for {
				fmt.Print("Введи A: ")
				var s string
				fmt.Scan(&s)
				a2, err2 := strconv.ParseFloat(s, 64)
				if err2 == nil {
					A = a2
					break
				}
				fmt.Println("Не число, давай ещё")
			}
		}
	} else {
		for {
			fmt.Print("Введи A: ")
			var s string
			fmt.Scan(&s)
			a2, err2 := strconv.ParseFloat(s, 64)
			if err2 == nil {
				A = a2
				break
			}
			fmt.Println("Не число, давай ещё")
		}
	}

	if len(args) > 1 {
		b, err := strconv.ParseFloat(args[1], 64)
		if err == nil {
			B = b
		} else {
			fmt.Println("Чёт B плохой, давай заново")
			for {
				fmt.Print("Введи B: ")
				var s string
				fmt.Scan(&s)
				b2, err2 := strconv.ParseFloat(s, 64)
				if err2 == nil {
					B = b2
					break
				}
				fmt.Println("Не число, давай ещё")
			}
		}
	} else {
		for {
			fmt.Print("Введи B: ")
			var s string
			fmt.Scan(&s)
			b2, err2 := strconv.ParseFloat(s, 64)
			if err2 == nil {
				B = b2
				break
			}
			fmt.Println("Не число, давай ещё")
		}
	}

	if len(args) > 2 {
		c, err := strconv.ParseFloat(args[2], 64)
		if err == nil {
			C = c
		} else {
			fmt.Println("Чёт C плохой, давай заново")
			for {
				fmt.Print("Введи C: ")
				var s string
				fmt.Scan(&s)
				c2, err2 := strconv.ParseFloat(s, 64)
				if err2 == nil {
					C = c2
					break
				}
				fmt.Println("Не число, давай ещё")
			}
		}
	} else {
		for {
			fmt.Print("Введи C: ")
			var s string
			fmt.Scan(&s)
			c2, err2 := strconv.ParseFloat(s, 64)
			if err2 == nil {
				C = c2
				break
			}
			fmt.Println("Не число, давай ещё")
		}
	}

	if A == 0 {
		fmt.Println("A ноль, это не квадратное, сорян")
		return
	}

	D := B*B - 4*A*C

	fmt.Println("Дискриминант =", D)

	if D > 0 {
		x1 := math.Sqrt((-B + math.Sqrt(D)) / (2 * A))
		x3 := -1 * x1
		x2 := math.Sqrt((-B - math.Sqrt(D)) / (2 * A))
		x4 := -1 * x2
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
		fmt.Println("x3 =", x3)
		fmt.Println("x4 =", x4)
	} else if D == 0 {
		x := -B / (2 * A)
		fmt.Println("x =", x)
	} else {
		fmt.Println("Корней нет, D отрицательный")
	}
}
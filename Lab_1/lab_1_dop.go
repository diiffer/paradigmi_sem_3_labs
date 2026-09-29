package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type BiquadraticEquation struct {
	A, B, C float64
}

func NewBiquadraticEquation(a, b, c float64) *BiquadraticEquation {
	return &BiquadraticEquation{A: a, B: b, C: c}
}

func (eq *BiquadraticEquation) Solve() []float64 {
	var roots []float64

	D := eq.B*eq.B - 4*eq.A*eq.C

	if D < 0 {
		return nil
	}

	var t1, t2 float64

	if D == 0 {
		t1 = -eq.B / (2 * eq.A)
		if t1 >= 0 {
			x := math.Sqrt(t1)
			roots = append(roots, x, -x)
		}
	} else {
		t1 = (-eq.B + math.Sqrt(D)) / (2 * eq.A)
		t2 = (-eq.B - math.Sqrt(D)) / (2 * eq.A)

		if t1 >= 0 {
			x := math.Sqrt(t1)
			roots = append(roots, x, -x)
		}
		if t2 >= 0 {
			x := math.Sqrt(t2)
			if len(roots) == 0 || math.Abs(x-roots[0]) > 1e-9 {
				roots = append(roots, x, -x)
			}
		}
	}

	return roots
}

type InputReader struct {
	reader *bufio.Reader
}

func NewInputReader() *InputReader {
	return &InputReader{reader: bufio.NewReader(os.Stdin)}
}

func (ir *InputReader) ReadCoefficient(name string) float64 {
	for {
		fmt.Printf("Введите коэффициент %s: ", name)
		input, _ := ir.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		val, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return val
		}
		fmt.Println("Некорректный ввод. Введите действительное число.")
	}
}

func (ir *InputReader) ReadAll() (float64, float64, float64) {
	a := ir.ReadCoefficient("A")
	b := ir.ReadCoefficient("B")
	c := ir.ReadCoefficient("C")
	return a, b, c
}

type Application struct {
	input *InputReader
}

func NewApplication() *Application {
	return &Application{input: NewInputReader()}
}

func (app *Application) GetCoefficients() (float64, float64, float64) {
	if len(os.Args) == 4 {
		a, errA := strconv.ParseFloat(os.Args[1], 64)
		b, errB := strconv.ParseFloat(os.Args[2], 64)
		c, errC := strconv.ParseFloat(os.Args[3], 64)

		if errA == nil && errB == nil && errC == nil {
			fmt.Println("Коэффициенты считаны из командной строки.")
			return a, b, c
		}
		fmt.Println("Ошибка в параметрах командной строки. Переход к ручному вводу.")
	}

	return app.input.ReadAll()
}

func (app *Application) PrintResult(roots []float64) {
	if len(roots) == 0 {
		fmt.Println("Корней нет")
		return
	}
	fmt.Println("Корни уравнения:")
	for _, root := range roots {
		fmt.Printf("%.4f\n", root)
	}
}

func (app *Application) Run() {
	fmt.Println("--- Решение биквадратного уравнения ---")

	a, b, c := app.GetCoefficients()
	eq := NewBiquadraticEquation(a, b, c)
	roots := eq.Solve()
	app.PrintResult(roots)
}

func main() {
	app := NewApplication()
	app.Run()
}

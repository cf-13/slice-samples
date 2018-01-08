package main

import (
	"strings"
	"fmt"
	"github.com/nh038426/print-function-in-out/printio"
	"golang.org/x/tour/pic"
)

const (
	AnderScore = "_"
	X = "X"
	O = "O"
	Func1 = 1
	Func2 = 2
	Func3 = 3
)

func main() {
	printio.FuncIn(main)

	sliceWithMake()
	slicesOfslices()
	appendSlices()
	rangeSample()
	rangeSample2()
	showPicture()

	printio.FuncOut(&printio.NotReturnFunction{main})
}

func returnFunction(i int) func(x, y int) uint8 {
	switch i {
	case Func1:
		return func(x, y int) uint8 { return uint8((x + y) / 2) }
	case Func2:
		return func(x, y int) uint8 { return uint8(x * y) }
	case Func3:
		return func(x, y int) uint8 { return uint8(x^y) }
	default:
		//Do Nothing
	}
	return func(x, y int) uint8 { return 0 }
}

func showPicture() {
	printio.FuncIn(showPicture)

	pic.Show(func(dx, dy int) ([][]uint8) {
		pics := make([][]uint8, dy)
		for i := 0; i < dy; i++ {
			pics[i] = make([]uint8, dx)
		}

		for x := range pics {
			for y := range pics[x] {
				f := returnFunction(Func1)
				pics[x][y] = f(x, y)
			}
		}
		return pics
	})

	printio.FuncOut(&printio.NotReturnFunction{showPicture})
}

func rangeSample2() {
	printio.FuncIn(rangeSample2)

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	printio.FuncOut(&printio.NotReturnFunction{rangeSample2})
}

func rangeSample() {
	printio.FuncIn(rangeSample)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}
	for i, v := range pow {
		fmt.Printf("pow[%d] = %d\n", i, v)
	}

	printio.FuncOut(&printio.NotReturnFunction{rangeSample})
}

func appendSlices() {
	printio.FuncIn(appendSlices)

	var s []int
	printSlice("s1", s)

	s = append(s, 0)
	printSlice("s2", s)

	s = append(s, 1)
	printSlice("s3", s)

	s = append(s, 2, 3, 4)
	printSlice("s4", s)

	printio.FuncOut(&printio.NotReturnFunction{appendSlices})
}

func slicesOfslices() {
	printio.FuncIn(slicesOfslices)

	board := [][]string{
		[]string{AnderScore, AnderScore, AnderScore},
		[]string{AnderScore, AnderScore, AnderScore},
		[]string{AnderScore, AnderScore, AnderScore},
	}

	board[0][0] = X
	board[2][2] = O
	board[1][2] = X
	board[1][0] = O
	board[0][2] = X

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	printio.FuncOut(&printio.NotReturnFunction{slicesOfslices})
}

func sliceWithMake() {
	printio.FuncIn(sliceWithMake)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	printio.FuncOut(&printio.NotReturnFunction{sliceWithMake})
}

func printSlice(s string, x []int){
	fmt.Printf("%s len = %d cap = %d %v\n",
		s, len(x), cap(x), x)
}
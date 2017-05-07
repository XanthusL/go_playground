package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"time"
)

var grids = [9][9]int{}

// Sudoku resolver
func main() {
	/*
		o o o | o o o | o o o
		o o o | o o o | o o o
		o o o | o o o | o o o
		------|-------|------
		o o o | o o o | o o o
		o o o | o o o | o o o
		o o o | o o o | o o o
		------|-------|------
		o o o | o o o | o o o
		o o o | o o o | o o o
		o o o | o o o | o o o
	*/
	checkFunc := func(row [9]int) {
		if e := check(row); e != nil {
			fmt.Println(e.Error())
			os.Exit(1)
		}
	}
	for j := 0; j < 9; j++ {
		fmt.Println("Type in the line", j+1, ":")
		// Input line by line
		fmt.Scanf("%d %d %d %d %d %d %d %d %d\n",
			&grids[j][0],
			&grids[j][1],
			&grids[j][2],
			&grids[j][3],
			&grids[j][4],
			&grids[j][5],
			&grids[j][6],
			&grids[j][7],
			&grids[j][8])
		// Row check
		check(grids[j])
	}

	// Column check
	traversalRow(checkFunc)
	// Sub grid check
	traversalSubGrid(checkFunc)
	// Initialize candidate data
	candidates := make(map[int][9]int)
	for i := 0; i < 81; i++ {
		candidates[i] = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	blankCount := 0
	traversalRow(func(r [9]int) {
		for _, v := range r {
			if v == 0 {
				blankCount++
			}
		}
	})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Too slow, there must be something wrong.")
		os.Exit(2)
	}()
	// Exclude numbers in candidates
	// Put left candidates in the grid.
	for blankCount > 0 {
		index := 0
		for i := range grids {
			for j, n := range grids[i] {
				if n != 0 {
					continue
				}
				numbers := getSubGrid(i, j)
				for _, sgv := range numbers {
					if sgv != 0 {
						if c, ok := candidates[index]; !ok {
							panic("Something wrong with candidates.")
						} else {
							c[sgv-1] = 0
						}
					}
				}
				numbers = getColumn(j)
				for _, sgv := range numbers {
					if sgv != 0 {
						if c, ok := candidates[index]; !ok {
							panic("Something wrong with candidates.")
						} else {
							c[sgv-1] = 0
						}
					}
				}
				numbers = grids[i]
				for _, sgv := range numbers {
					if sgv != 0 {
						if c, ok := candidates[index]; !ok {
							panic("Something wrong with candidates.")
						} else {
							c[sgv-1] = 0
						}
					}
				}
				tmp := 0
				for _, c := range candidates[index] {
					if c != 0 {
						if tmp == 0 {
							tmp = c
						} else {
							tmp = 0
							break
						}
					}
				}
				grids[i][j] = tmp
				if tmp != 0 {
					blankCount--
				}
				index++
				panic("It doesn't work, the algorithm is wrong.")
			}
		}
	}
	fmt.Println(grids)
}
func check(numbers [9]int) error {
	ns := make([]int, 9)
	copy(ns, numbers[:])
	sort.Ints(ns)
	for i, v := range ns {
		if i == 0 || v == 0 {
			continue
		}
		if v == ns[i-1] {
			return errors.New(fmt.Sprintf("Error: Numbers can't repeat ,%v", numbers))
		}
		if v > 9 {
			return errors.New("Error : Number range is [0,9]")
		}
	}
	return nil
}

func traversalRow(f func([9]int)) {
	for i := 0; i < 9; i++ {
		c := [...]int{
			grids[0][i],
			grids[1][i],
			grids[2][i],
			grids[3][i],
			grids[4][i],
			grids[5][i],
			grids[6][i],
			grids[7][i],
			grids[8][i],
		}
		f(c)
	}
}
func traversalSubGrid(f func([9]int)) {
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			x := 3 * j
			y := 3 * k
			numbers := [...]int{
				grids[x][y],
				grids[x][y+1],
				grids[x][y+2],
				grids[x+1][y],
				grids[x+1][y+1],
				grids[x+1][y+2],
				grids[x+2][y],
				grids[x+2][y+1],
				grids[x+2][y+2]}
			f(numbers)
		}
	}

}
func getSubGrid(x, y int) [9]int {
	x /= 3
	x *= 3
	y /= 3
	y *= 3
	return [...]int{
		grids[x][y],
		grids[x][y+1],
		grids[x][y+2],
		grids[x+1][y],
		grids[x+1][y+1],
		grids[x+1][y+2],
		grids[x+2][y],
		grids[x+2][y+1],
		grids[x+2][y+2]}
}
func getColumn(i int) [9]int {
	return [...]int{
		grids[0][i],
		grids[1][i],
		grids[2][i],
		grids[3][i],
		grids[4][i],
		grids[5][i],
		grids[6][i],
		grids[7][i],
		grids[8][i],
	}
}

/*
Test case:

Type in the line 1 :
0 0 4 3 0 2 6 0 0
Type in the line 2 :
0 0 8 0 0 0 9 0 0
Type in the line 3 :
0 3 0 0 7 0 0 2 0
Type in the line 4 :
0 0 3 0 0 0 2 0 0
Type in the line 5 :
9 0 0 0 1 0 0 0 8
Type in the line 6 :
4 0 0 7 0 8 0 0 1
Type in the line 7 :
0 0 7 9 0 5 1 0 0
Type in the line 8 :
0 5 0 8 2 1 0 3 0
Type in the line 9 :
0 0 0 0 0 0 0 0 0

*/

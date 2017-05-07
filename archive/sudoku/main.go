package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
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
		if e := check(grids[j]); e != nil {
			fmt.Println(e.Error())
			return
		}

	}
	// Column check
	//for i := 0; i < 9; i++ {
	//	c := [...]int{
	//		grids[0][i],
	//		grids[1][i],
	//		grids[2][i],
	//		grids[3][i],
	//		grids[4][i],
	//		grids[5][i],
	//		grids[6][i],
	//		grids[7][i],
	//		grids[8][i],
	//	}
	//	if e := check(c); e != nil {
	//		fmt.Println(e.Error())
	//		return
	//	}
	//}
	// Column check
	traversalRow(func(row [9]int) {
		if e := check(row); e != nil {
			fmt.Println(e.Error())
			os.Exit(1)
		}
	})
	// TODO : Sub grid check

	// Initialize candidate data
	candidates := make(map[int][9]int)
	for i := 0; i < 81; i++ {
		candidates[i] = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	// TODO : Exclude numbers in candidates

	// TODO : Put left candidates in the grid.

	// TODO : Optimize,  functions to traversal column and sub-grid

	fmt.Println(grids)
}
func check(numbers [9]int) error {
	ns := make([]int, 9)
	copy(ns, numbers[:])
	sort.Ints(ns)
	for i, v := range ns {
		if i == 0 {
			continue
		}
		if v == ns[i-1] {
			return errors.New("Error: Numbers can't repeat")
		}
		if v > 9 {
			return errors.New("ERROR : Number range is [0,9]")
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

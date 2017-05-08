package main

import (
	"errors"
	"fmt"
	"lyh_demos/golang/mylib"
	"os"
	"sort"
	"time"
)

var grids = [9][9]int{}
var blankCount = 0
var candidates = make(map[int]*mylib.IntSet)

// Sudoku resolver
func main() {
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
		check(grids[j], false)
	}
	checkFunc := func(row [9]int) {
		if e := check(row, false); e != nil {
			fmt.Println(e.Error())
			os.Exit(1)
		}
	}
	// Column check
	traversalRow(checkFunc)
	// Sub grid check
	traversalSubGrid(checkFunc)
	// Initialize candidate data
	for i := 0; i < 81; i++ {
		s := &mylib.IntSet{}
		s.Init()
		s.Add(1)
		s.Add(2)
		s.Add(3)
		s.Add(4)
		s.Add(5)
		s.Add(6)
		s.Add(7)
		s.Add(8)
		s.Add(9)
		candidates[i] = s
	}
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
	exclude()
	// Try one by one
	// tryingTable := make(map[int]int)
	backup := [81]int{}
	for i := range grids {
		for j, n := range grids[i] {
			if n == 0 {
				bigIndex := i*9 + j
				m := candidates[bigIndex].GetMap()
				if len(m) == 0 {
					panic("No candidates to use.")
				}
				for c := range m {
					backup[bigIndex] = c
					break
				}
			}
		}
	}
	//
	ok := true
	validate := func(r [9]int) {
		e := check(r, true)
		if e != nil {
			ok = false
		}
	}
	traversalRow(validate)
	traversalSubGrid(validate)
	for i := range grids {
		validate(grids[i])
	}
	if ok {
		fmt.Println(grids)
		return
	} else {
		for i := range backup {
			grids[i/9][i%9] = 0
			backup[i] = 0
		}
	}

}
func check(numbers [9]int, full bool) error {
	ns := make([]int, 9)
	copy(ns, numbers[:])
	sort.Ints(ns)
	for i, v := range ns {
		if i == 0 {
			continue
		}
		if !full && v == 0 {
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
func exclude() error {
	for i := range grids {
		for j, n := range grids[i] {
			if n != 0 {
				continue
			}
			exclude := func(n [9]int) {
				for _, v := range n {
					if v != 0 {
						if c, ok := candidates[9*i+j]; !ok {
							panic("Something wrong with candidates.")
						} else {
							c.Remove(v)
						}
					}
				}
			}
			numbers := getSubGrid(i, j)
			exclude(numbers)
			numbers = getColumn(j)
			exclude(numbers)
			numbers = grids[i]
			exclude(numbers)
			if candidates[9*i+j].Size() == 1 {
				m := candidates[9*i+j].GetMap()
				if len(m) == 0 {
					return errors.New("No candidates can use.")
				}
				for N := range m {
					grids[i][j] = N
					blankCount--
				}
			}
		}
	}
	return nil
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

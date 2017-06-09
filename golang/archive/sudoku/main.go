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
var candidates = make(map[int]*mylib.IntSet)

type Node struct {
	Options     []int
	IndexInGrid int
	Next        *Node
}

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

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Too slow, there must be something wrong.")
		fmt.Println(grids)
		os.Exit(2)
	}()
	// Exclude numbers in candidates
	exclude()
	// Try one by one
	// List of options for blanks
	var solution *Node
	var nodeCursor *Node
	for i := range grids {
		for j, n := range grids[i] {
			bigIndex := i*9 + j
			if n == 0 {
				s := candidates[bigIndex]
				n := &Node{
					IndexInGrid: bigIndex,
					Options:     s.ToSlice(),
				}
				if solution == nil {
					solution = n
				} else {
					nodeCursor.Next = n
				}
				nodeCursor = n
			}
		}
	}
	fmt.Println(walk(solution))
	fmt.Println(grids)

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
		f(grids[i])
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
	exclude := func(i int, n [9]int) {
		for _, v := range n {
			if v != 0 {
				if c, ok := candidates[i]; !ok {
					panic("Something wrong with candidates.")
				} else {
					c.Remove(v)
				}
			}
		}
	}
	for i := range grids {
		for j, n := range grids[i] {
			if n != 0 {
				delete(candidates, i*9+j)
				continue
			}
			numbers := getSubGrid(i, j)

			exclude(i*9+j, numbers)
			numbers = getColumn(j)
			exclude(i*9+j, numbers)
			numbers = grids[i]
			exclude(i*9+j, numbers)
			if candidates[9*i+j].Size() == 1 {
				m := candidates[9*i+j].GetMap()
				if len(m) == 0 {
					return errors.New("No candidates can use.")
				}
				for N := range m {
					grids[i][j] = N
				}
				delete(candidates, i*9+j)
			} else {
				fmt.Println(candidates[9*i+j].ToSlice())
			}
		}
	}
	return nil
}
func walk(n *Node) bool {
	if n == nil {
		return isValidate()
	}
trying:
	for _, v := range n.Options {
		// Get all values related to grids[n.IndexInGrid/9][n.IndexInGrid%9]
		r := grids[n.IndexInGrid%9]
		c := getColumn(n.IndexInGrid / 9)
		g := getSubGrid(n.IndexInGrid/9, n.IndexInGrid%9)
		rlt := append(r[:], c[:]...)
		rlt = append(rlt, g[:]...)

		// if unusable, continue
		// else try next blank
		for _, rv := range rlt {
			if rv != 0 && rv == v {
				continue trying
			}
		}

		grids[n.IndexInGrid/9][n.IndexInGrid%9] = v
		if walk(n.Next) {
			return true
		} else {
			// Cleanup value of the blank
			grids[n.IndexInGrid/9][n.IndexInGrid%9] = 0
		}
	}
	return false
}
func isValidate() bool {
	for i := 0; i < 9; i++ {
		c := getColumn(i)
		e := check(c, true)
		if e != nil {
			return false
		}
		c = grids[i]
		e = check(c, true)
		if e != nil {
			return false
		}
	}
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			x := 3 * j
			y := 3 * k
			g := getSubGrid(x, y)
			e := check(g, true)
			if e != nil {
				return false
			}
		}
	}
	return true
}

/*
Test case:

0 0 0 0 0 6 0 9 0
0 8 0 5 0 0 7 0 0
0 4 9 0 0 0 5 0 0
0 3 0 8 1 0 0 0 5
4 0 0 0 0 0 0 0 9
2 0 0 0 9 7 0 8 0
0 0 5 0 0 0 2 1 0
0 0 6 0 0 5 0 4 0
0 7 0 3 0 0 0 0 0


*/

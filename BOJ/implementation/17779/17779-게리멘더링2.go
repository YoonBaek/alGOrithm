// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var (
// 	n, total int
// 	diffs    []int
// 	district [][]int
// 	test     [][]string
// 	sc       = bufio.NewScanner(os.Stdin)
// )

// func init() {
// 	sc.Split(bufio.ScanWords)
// }

// func main() {
// 	n = scan()
// 	district = make([][]int, n)
// 	test = make([][]string, n)
// 	for row := range district {
// 		district[row] = make([]int, n)
// 		test[row] = make([]string, n)
// 		for col := range district[row] {
// 			district[row][col] = scan()
// 			test[row][col] = "0"
// 			total += district[row][col]
// 		}
// 	}
// 	for row := 0; row < n-1; row++ {
// 		for col := 0; col < n-1; col++ {
// 			gm(row, col)
// 		}
// 	}
// 	fmt.Println(min(diffs...), district)
// }

// func scan() int {
// 	sc.Scan()
// 	n, _ := strconv.Atoi(sc.Text())
// 	return n
// }

// func gm(row, col int) {
// 	for d1 := 1; d1 <= col; d1++ {
// 		for d2 := 1; d2 <= n-1-col; d2++ {
// 			var one, two, three, four, five int
// 			for r := range district {
// 				for c := range district {
// 					if r < row+d1 && c <= col && c < (col-(r-row)) {
// 						test[r][c] = "1"
// 						one += district[r][c]
// 					}
// 					if r <= row+d2 && c > col && c > (col+(r-row)) {
// 						test[r][c] = "2"
// 						two += district[r][c]
// 					}
// 					if r >= row+d1 && c < col-d1+d2 && c < (col-d1+(r-(row+d1))) {
// 						test[r][c] = "3"
// 						three += district[r][c]
// 					}
// 					if r > row+d2 && c >= col-d1+d2 && c > (col+d2-(r-(row+d2))) {
// 						test[r][c] = "4"
// 						four += district[r][c]
// 					}
// 				}
// 			}
// 			five = total - one - two - three - four
// 			fmt.Println(row, col, d1, d2)
// 			fmt.Println(one, two, three, four, five)
// 			for row := range test {
// 				fmt.Println(test[row])
// 				test[row] = make([]string, n)
// 				for col := range test[row] {
// 					test[row][col] = "0"
// 				}
// 			}
// 			diff := max(one, two, three, four, five) - min(one, two, three, four, five)
// 			diffs = append(diffs, diff)
// 		}
// 	}
// }

// func max(nums ...int) (max int) {
// 	for _, num := range nums {
// 		if num > max {
// 			max = num
// 		}
// 	}
// 	return
// }

// func min(nums ...int) (min int) {
// 	min = 20 * 20 * 100
// 	for _, num := range nums {
// 		if num < min {
// 			min = num
// 		}
// 	}
// 	return
// }
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n, total int
	diffs    []int
	district [][]int
	sc       = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n = scan()
	district = make([][]int, n)
	for row := range district {
		district[row] = make([]int, n)
		for col := range district[row] {
			district[row][col] = scan()
			total += district[row][col]
		}
	}
	for row := 0; row < n-1; row++ {
		for col := 0; col < n-1; col++ {
			gm(row, col)
		}
	}
	fmt.Println(min(diffs...))
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func gm(row, col int) {
	for d1 := 1; d1 <= col; d1++ {
		for d2 := 1; d2 <= n-1-col; d2++ {
			var one, two, three, four, five int

			for r := range district {
				for c := range district {
					if r < row+d1 && c <= col && c < (col-(r-row)) {
						one += district[r][c]
					}
					if r <= row+d2 && c > col && c > (col+(r-row)) {
						two += district[r][c]
					}
					if r >= row+d1 && c < col-d1+d2 && c < (col-d1+(r-(row+d1))) {
						three += district[r][c]
					}
					if r > row+d2 && c >= col-d1+d2 && c > (col+d2-(r-(row+d2))) {
						four += district[r][c]
					}
				}
			}

			five = total - one - two - three - four
			diff := max(one, two, three, four, five) - min(one, two, three, four, five)
			diffs = append(diffs, diff)
		}
	}
}

func max(nums ...int) (max int) {
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return
}

func min(nums ...int) (min int) {
	min = 20 * 20 * 100
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return
}

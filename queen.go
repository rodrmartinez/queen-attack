package queen

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */

func get_up(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var c [][]int32
	for j := r_q; j <= n; j++ {
		if j != r_q {
			for _, o := range obstacles {
				if j == o[0] && c_q == o[1] {
					return c
				}
			}
			c = append(c, []int32{int32(j), c_q})
		}
	}
	return c
}

func get_down(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var c [][]int32
	for j := r_q; j >= 1; j-- {
		if j != r_q {
			for _, o := range obstacles {
				if j == o[0] && c_q == o[1] {
					return c
				}
			}
			c = append(c, []int32{int32(j), c_q})
		}
	}
	return c
}

func get_right(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var r [][]int32
	for j := c_q; j <= n; j++ {
		if j != c_q {
			for _, o := range obstacles {
				if r_q == o[0] && j == o[1] {
					return r
				}
			}
			r = append(r, []int32{r_q, int32(j)})
		}
	}
	return r
}

func get_left(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var r [][]int32
	for j := c_q; j >= 1; j-- {
		if j != c_q {
			for _, o := range obstacles {
				if r_q == o[0] && j == o[1] {
					return r
				}
			}
			r = append(r, []int32{r_q, int32(j)})
		}
	}
	return r
}

func get_diag_ur(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var d [][]int32

	//UP RIGHT
	row := r_q
	for j := c_q; j <= n; j++ {
		if j != c_q && row < n {
			row++
			for _, o := range obstacles {
				if row == o[0] && j == o[1] {
					return d
				}
			}
			d = append(d, []int32{int32(row), int32(j)})
		}
	}
	return d
}

func get_diag_dr(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var d [][]int32

	// DOWN RIGHT
	row := r_q
	for j := c_q; j <= n; j++ {
		if j != c_q && row > 1 {
			row--
			for _, o := range obstacles {
				if row == o[0] && j == o[1] {
					return d
				}
			}
			d = append(d, []int32{int32(row), int32(j)})
		}
	}
	return d
}
func get_diag_ul(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var d [][]int32

	// UP LEFT
	col := c_q
	for j := r_q; j <= n; j++ {
		if j != r_q && col > 1 {
			col--
			for _, o := range obstacles {
				if j == o[0] && col == o[1] {
					return d
				}
			}
			d = append(d, []int32{int32(j), int32(col)})
		}
	}
	return d
}

func get_diag_dl(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) [][]int32 {
	var d [][]int32

	//DOWN LEFT
	col := c_q
	for j := r_q; j >= 1; j-- {
		if j != r_q && col > 1 {
			col--
			for _, o := range obstacles {
				if j == o[0] && col == o[1] {
					return d
				}
			}
			d = append(d, []int32{int32(j), int32(col)})
		}
	}
	return d
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var squares [][]int32

	// // //column squares
	up := get_up(n, k, r_q, c_q, obstacles)
	down := get_down(n, k, r_q, c_q, obstacles)
	left := get_left(n, k, r_q, c_q, obstacles)
	right := get_right(n, k, r_q, c_q, obstacles)
	diag_ur := get_diag_ur(n, k, r_q, c_q, obstacles)
	diag_dr := get_diag_dr(n, k, r_q, c_q, obstacles)
	diag_ul := get_diag_ul(n, k, r_q, c_q, obstacles)
	diag_dl := get_diag_dl(n, k, r_q, c_q, obstacles)

	squares = append(squares, up...)
	squares = append(squares, down...)
	squares = append(squares, left...)
	squares = append(squares, right...)
	squares = append(squares, diag_ur...)
	squares = append(squares, diag_dr...)
	squares = append(squares, diag_ul...)
	squares = append(squares, diag_dl...)

	return int32(len(squares))
}

package main

import (
	"math"
)

type Matrix struct {
	arr [][]float64
	n   int
	m   int
}

func (a *Matrix) init(n, m int) {
	a.arr = make([][]float64, n)
	for i := range a.arr {
		a.arr[i] = make([]float64, m)
	}
	a.n = n
	a.m = m
}

func (a Matrix) mult(b Matrix) Matrix {
	n := a.n
	k := a.m
	m := b.m
	answer := Matrix{nil, n, m}
	answer.init(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for p := 0; p < m; p++ {

				answer.arr[i][p] += a.arr[i][j] * b.arr[j][p]
			}
		}
	}
	return answer
}

func (a Matrix) transpose() Matrix {
	n := a.n
	m := a.m
	answer := Matrix{nil, 0, 0}
	answer.init(m, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			answer.arr[j][i] = a.arr[i][j]
		}
	}
	return answer
}

func gauss(a Matrix) Matrix {
	const EPS = 1e-9
	where := make([]int, a.m)
	for i := 0; i < a.m; i++ {
		where[i] = -1
	}
	row := 0
	m := a.m / 2
	for col := 0; col < m; col++ {
		if row >= a.n {
			break
		}
		sel := row
		for i := row; i < a.n; i++ {
			if math.Abs(a.arr[i][col]) > math.Abs(a.arr[sel][col]) {
				sel = i
			}
		}
		if math.Abs(a.arr[sel][col]) < EPS {
			continue
		}
		for i := col; i < a.m; i++ {
			a.arr[sel][i], a.arr[row][i] = a.arr[row][i], a.arr[sel][i]
		}
		where[col] = row
		for i := 0; i < a.n; i++ {
			if i != row {
				c := a.arr[i][col] / a.arr[row][col]
				for j := col; j < a.m; j++ {
					a.arr[i][j] -= a.arr[row][j] * c
				}
			}
		}
		row++
	}
	ans := Matrix{nil, a.n, a.n}
	ans.init(a.n, a.n)
	for i := 0; i < m; i++ {
		if where[i] != -1 {
			for j := m; j < a.m; j++ {
				ans.arr[i][j-m] = a.arr[where[i]][j] / a.arr[where[i]][i]
			}
		}
	}
	return ans
}

func (a Matrix) inversed() Matrix {

	b := Matrix{nil, a.n, a.n * 2}
	b.init(a.n, a.n*2)
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			b.arr[i][j] = a.arr[i][j]
		}
		b.arr[i][i+a.n] = 1
	}
	gauss_solution := gauss(b)
	return gauss_solution
}

func (a Matrix) equals(b Matrix) bool {
	const EPS = 1e-9
	if a.n != b.n || a.m != b.m {
		return false
	}
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			if math.Abs(a.arr[i][j]-b.arr[i][j]) > EPS {
				return false
			}
		}
	}
	return true
}

func eye(n int) Matrix {
	ans := Matrix{nil, n, n}
	ans.init(n, n)
	for i := 0; i < n; i++ {
		ans.arr[i][i] = 1
	}
	return ans
}

func (a Matrix) add(b Matrix) Matrix {
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			a.arr[i][j] += b.arr[i][j]
		}
	}
	return a
}

func (a Matrix) scalarMult(alpha float64) Matrix {
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			a.arr[i][j] *= alpha
		}
	}
	return a
}

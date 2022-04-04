package src

import (
	"math"
)

type Matrix struct {
	Arr [][]float64
	N   int
	M   int
}

func (a *Matrix) init(n, m int) {
	a.Arr = make([][]float64, n)
	for i := range a.Arr {
		a.Arr[i] = make([]float64, m)
	}
	a.N = n
	a.M = m
}

func (a Matrix) Mult(b Matrix) Matrix {
	n := a.N
	k := a.M
	m := b.M
	answer := Matrix{nil, n, m}
	answer.init(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for p := 0; p < m; p++ {

				answer.Arr[i][p] += a.Arr[i][j] * b.Arr[j][p]
			}
		}
	}
	return answer
}

func (a Matrix) Transpose() Matrix {
	n := a.N
	m := a.M
	answer := Matrix{nil, 0, 0}
	answer.init(m, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			answer.Arr[j][i] = a.Arr[i][j]
		}
	}
	return answer
}

func gauss(a Matrix) Matrix {
	const EPS = 1e-9
	where := make([]int, a.M)
	for i := 0; i < a.M; i++ {
		where[i] = -1
	}
	row := 0
	m := a.M / 2
	for col := 0; col < m; col++ {
		if row >= a.N {
			break
		}
		sel := row
		for i := row; i < a.N; i++ {
			if math.Abs(a.Arr[i][col]) > math.Abs(a.Arr[sel][col]) {
				sel = i
			}
		}
		if math.Abs(a.Arr[sel][col]) < EPS {
			continue
		}
		for i := col; i < a.M; i++ {
			a.Arr[sel][i], a.Arr[row][i] = a.Arr[row][i], a.Arr[sel][i]
		}
		where[col] = row
		for i := 0; i < a.N; i++ {
			if i != row {
				c := a.Arr[i][col] / a.Arr[row][col]
				for j := col; j < a.M; j++ {
					a.Arr[i][j] -= a.Arr[row][j] * c
				}
			}
		}
		row++
	}
	ans := Matrix{nil, a.N, a.N}
	ans.init(a.N, a.N)
	for i := 0; i < m; i++ {
		if where[i] != -1 {
			for j := m; j < a.M; j++ {
				ans.Arr[i][j-m] = a.Arr[where[i]][j] / a.Arr[where[i]][i]
			}
		}
	}
	return ans
}

func (a Matrix) Inversed() Matrix {

	b := Matrix{nil, a.N, a.N * 2}
	b.init(a.N, a.N*2)
	for i := 0; i < a.N; i++ {
		for j := 0; j < a.N; j++ {
			b.Arr[i][j] = a.Arr[i][j]
		}
		b.Arr[i][i+a.N] = 1
	}
	gauss_solution := gauss(b)
	return gauss_solution
}

func (a Matrix) Equals(b Matrix) bool {
	const EPS = 1e-9
	if a.N != b.N || a.M != b.M {
		return false
	}
	for i := 0; i < a.N; i++ {
		for j := 0; j < a.M; j++ {
			if math.Abs(a.Arr[i][j]-b.Arr[i][j]) > EPS {
				return false
			}
		}
	}
	return true
}

func Eye(n int) Matrix {
	ans := Matrix{nil, n, n}
	ans.init(n, n)
	for i := 0; i < n; i++ {
		ans.Arr[i][i] = 1
	}
	return ans
}

func (a Matrix) add(b Matrix) Matrix {
	for i := 0; i < a.N; i++ {
		for j := 0; j < a.M; j++ {
			a.Arr[i][j] += b.Arr[i][j]
		}
	}
	return a
}

func (a Matrix) scalarMult(alpha float64) Matrix {
	for i := 0; i < a.N; i++ {
		for j := 0; j < a.M; j++ {
			a.Arr[i][j] *= alpha
		}
	}
	return a
}

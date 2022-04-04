package tests

import (
	"awesomeProject1/src"
	"fmt"
	"testing"
)

func TestMatrix_inversed(t *testing.T) {
	a := src.Matrix{[][]float64{{5.0, 5.0, 5.0}, {5.0, 5.0, 6.0}, {7.0, 8.0, 7.0}}, 3, 3}
	inversedA := a.Inversed()
	result := a.Mult(inversedA)
	E := src.Eye(3)
	fmt.Println(result)
	if !result.Equals(E) {
		t.Error("incorrect matrix inversion")
	}
}

func TestMatrix_mult(t *testing.T) {
	a := src.Matrix{[][]float64{{1.0, 2.0}, {3.0, 4.0}, {5.0, 6.0}}, 3, 2}
	b := src.Matrix{[][]float64{{5.0}, {2.0}}, 2, 1}
	expected_result := src.Matrix{[][]float64{{9.0}, {23.0}, {37.0}}, 3, 1}
	if !a.Mult(b).Equals(expected_result) {
		t.Error("incorrect matrix multiplication")
	}
}

func TestMatrix_transpose(t *testing.T) {
	a := src.Matrix{[][]float64{{1, 2}, {3, 4}, {5, 6}}, 3, 2}
	aT := src.Matrix{[][]float64{{1, 3, 5}, {2, 4, 6}}, 2, 3}
	if !a.Transpose().Equals(aT) {
		t.Error("incorrect matrix transponion")
	}

}

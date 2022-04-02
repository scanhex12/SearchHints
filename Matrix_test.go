package main

import (
	"fmt"
	"testing"
)

func TestMatrix_inversed(t *testing.T) {
	a := Matrix{[][]float64{{5.0, 5.0, 5.0}, {5.0, 5.0, 6.0}, {7.0, 8.0, 7.0}}, 3, 3}
	inversedA := a.inversed()
	result := a.mult(inversedA)
	E := eye(3)
	fmt.Println(result)
	if !result.equals(E) {
		t.Error("incorrect matrix inversion")
	}
}

func TestMatrix_mult(t *testing.T) {
	a := Matrix{[][]float64{{1.0, 2.0}, {3.0, 4.0}, {5.0, 6.0}}, 3, 2}
	b := Matrix{[][]float64{{5.0}, {2.0}}, 2, 1}
	expected_result := Matrix{[][]float64{{9.0}, {23.0}, {37.0}}, 3, 1}
	if !a.mult(b).equals(expected_result) {
		t.Error("incorrect matrix multiplication")
	}
}

func TestMatrix_transpose(t *testing.T) {
	a := Matrix{[][]float64{{1, 2}, {3, 4}, {5, 6}}, 3, 2}
	aT := Matrix{[][]float64{{1, 3, 5}, {2, 4, 6}}, 2, 3}
	if !a.transpose().equals(aT) {
		t.Error("incorrect matrix transponion")
	}

}

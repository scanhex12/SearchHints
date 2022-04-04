package tests

import (
	"awesomeProject1/src"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := src.Trie{}
	tree.Fit([][]string{{"VK", "is", "a", "good", "company"},
		{"Rock", "am", "using", "VK"},
		{"Rocket", "was", "invented"}})
}

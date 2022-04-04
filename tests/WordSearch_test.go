package tests

import (
	"awesomeProject1/src"
	"testing"
)

func TestIntersectSortedLists(t *testing.T) {
	list1 := []int{1, 3, 4, 7, 9, 11, 13, 24}
	list2 := []int{0, 2, 3, 5, 7, 9, 15, 24, 125}
	answer := src.IntersectSortedLists(list1, list2)
	expected_answer := []int{3, 7, 9, 24}
	if len(answer) != len(expected_answer) {
		t.Error("Incorrect answer for list intersection")
	}
	for i := 0; i < len(answer); i++ {
		if answer[i] != expected_answer[i] {
			t.Error("Incorrect answer for list intersection")
		}
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestWordSearch(t *testing.T) {
	model := src.WordSearch{}
	sentences := [][]string{{"it", "is", "vk"},
		{"it", "is", "the", "best", "company"}}

	model.Fit(sentences)
	if len(model.PhraseCount) != 5 ||
		model.PhraseCount[src.MakePair("it", "is")] != 2 ||
		model.PhraseCount[src.MakePair("is", "vk")] != 1 ||
		model.PhraseCount[src.MakePair("is", "the")] != 1 ||
		model.PhraseCount[src.MakePair("the", "best")] != 1 ||
		model.PhraseCount[src.MakePair("best", "company")] != 1 {
		t.Error("Incorrect PhraseCount values")
	}

	if len(model.AllWords) != 6 ||
		model.AllWords[0] != "it" ||
		model.AllWords[1] != "is" ||
		model.AllWords[2] != "vk" ||
		model.AllWords[3] != "the" ||
		model.AllWords[4] != "best" ||
		model.AllWords[5] != "company" {
		t.Error("Incorrect AllWords values")
	}

	if len(model.WordId) != 6 ||
		model.WordId["it"] != 1 ||
		model.WordId["is"] != 2 ||
		model.WordId["vk"] != 3 ||
		model.WordId["the"] != 4 ||
		model.WordId["best"] != 5 ||
		model.WordId["company"] != 6 {
		t.Error("Incorrect WordId values")
	}

	if len(model.InvertedIndex) != 7 ||
		!IntArrayEquals(model.InvertedIndex[1], []int{0, 1}) ||
		!IntArrayEquals(model.InvertedIndex[2], []int{0, 1}) ||
		!IntArrayEquals(model.InvertedIndex[3], []int{0}) ||
		!IntArrayEquals(model.InvertedIndex[4], []int{1}) ||
		!IntArrayEquals(model.InvertedIndex[5], []int{1}) ||
		!IntArrayEquals(model.InvertedIndex[6], []int{1}) {
		t.Error("Incorrect InvertedIndex values")
	}

	if len(model.Sentences) != 2 ||
		!IntArrayEquals(model.Sentences[0], []int{1, 2, 3}) ||
		!IntArrayEquals(model.Sentences[1], []int{1, 2, 4, 5, 6}) {
		t.Error("Incorrect Sentences values")

	}
}

func TestImportantIndexes(t *testing.T) {
	model := src.WordSearch{}
	sentences := [][]string{{"it", "is", "vk"},
		{"it", "is", "the", "best", "company"},
		{"the", "company"},
		{"the"}}
	model.Fit(sentences)
	query := []string{"vk", "the", "it", "best"}
	expected_answer := []int{0, 3}
	if !IntArrayEquals(model.GetImportantWords(query), expected_answer) {
		t.Error("Incorrect important indexes")
	}

	query = []string{"vk", "query", "the", "it"}
	expected_answer = []int{0, 1, 3}
	if !IntArrayEquals(model.GetImportantWords(query), expected_answer) {
		t.Error("Incorrect important indexes")
	}

}

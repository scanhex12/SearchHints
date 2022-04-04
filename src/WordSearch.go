package src

import "sort"

type Pair struct {
	Values [2]interface{}
}

func MakePair(k, v interface{}) Pair {
	return Pair{Values: [2]interface{}{k, v}}
}

func (p Pair) Get(i int) interface{} {
	return p.Values[i]
}

type WordSearch struct {
	tfidf         *TfIdf
	PhraseCount   map[Pair]int
	WordId        map[string]int
	AllWords      []string
	InvertedIndex [][]int
	Sentences     [][]int
}

func IntersectSortedLists(list1, list2 []int) []int {
	it2 := 0
	answer := make([]int, 0)
	for it1 := 0; it1 < len(list1); it1++ {
		for list2[it2] < list1[it1] && it2 < len(list2)-1 {
			it2++
		}
		if list1[it1] == list2[it2] {
			answer = append(answer, list1[it1])
		}
	}
	return answer
}

func (model *WordSearch) Fit(trainSentences [][]string) {
	model.tfidf = &TfIdf{}
	model.PhraseCount = make(map[Pair]int)
	model.WordId = make(map[string]int)
	model.AllWords = make([]string, 0)
	model.InvertedIndex = make([][]int, 0)
	model.Sentences = make([][]int, len(trainSentences))

	model.tfidf.Fit(trainSentences)

	for i := 0; i < len(trainSentences); i++ {
		for j := 0; j < len(trainSentences[i])-1; j++ {
			model.PhraseCount[MakePair(trainSentences[i][j], trainSentences[i][j+1])] += 1
		}
	}
	for i := 0; i < len(trainSentences); i++ {
		for j := 0; j < len(trainSentences[i]); j++ {
			if model.WordId[trainSentences[i][j]] == 0 {
				model.AllWords = append(model.AllWords, trainSentences[i][j])
				model.WordId[trainSentences[i][j]] = len(model.AllWords)
			}
			model.Sentences[i] = append(model.Sentences[i], model.WordId[trainSentences[i][j]])
		}
	}
	model.InvertedIndex = make([][]int, len(model.AllWords)+1)
	for i := 0; i < len(trainSentences); i++ {
		for j := 0; j < len(trainSentences[i]); j++ {
			if len(model.InvertedIndex[model.WordId[trainSentences[i][j]]]) == 0 {
				model.InvertedIndex[model.WordId[trainSentences[i][j]]] =
					append(model.InvertedIndex[model.WordId[trainSentences[i][j]]], i)
				continue
			}
			lastIndex := len(model.InvertedIndex[model.WordId[trainSentences[i][j]]]) - 1
			if model.InvertedIndex[model.WordId[trainSentences[i][j]]][lastIndex] == i {
				continue
			}
			model.InvertedIndex[model.WordId[trainSentences[i][j]]] =
				append(model.InvertedIndex[model.WordId[trainSentences[i][j]]], i)
		}
	}
}

func (model *WordSearch) GetImportantWords(query []string) []int {
	if len(query) == 0 {
		return make([]int, 0)
	}
	if len(query) == 1 {
		return make([]int, 1)
	}
	tfidfs := model.tfidf.TfIdfAllSentence(query)
	tfidfsCopy := make([]float64, len(tfidfs))
	copy(tfidfsCopy, tfidfs)
	sort.Float64s(tfidfs)
	importantIndexes := []int{-1, -1}

	for i := 0; i < len(query); i++ {
		if tfidfsCopy[i] == tfidfs[0] && importantIndexes[0] == -1 {
			importantIndexes[0] = i
		}
		if tfidfsCopy[i] == tfidfs[1] && importantIndexes[1] == -1 {
			importantIndexes[1] = i
		}
	}

	if len(query) >= 3 {
		if len(query)-1 != importantIndexes[0] && len(query)-1 != importantIndexes[1] {
			importantIndexes = append(importantIndexes, len(query)-1)
		}
	}

	return importantIndexes
}

func (model *WordSearch) Predict(query []string) string {
	return ""
}

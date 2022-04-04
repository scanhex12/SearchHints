package src

import "math"

type TfIdf struct {
	wordDocFreq map[string]int
	numDocs     int
}

func (tfidf *TfIdf) Fit(documents [][]string) {
	tfidf.wordDocFreq = make(map[string]int)
	numDocks := len(documents)
	for i := 0; i < numDocks; i++ {
		curDocVocab := make(map[string]bool)
		for j := 0; j < len(documents[i]); j++ {
			if !curDocVocab[documents[i][j]] {
				tfidf.wordDocFreq[documents[i][j]] += 1
			} else {
				curDocVocab[documents[i][j]] = true
			}
		}
	}
	tfidf.numDocs = numDocks
}

func (tfidf *TfIdf) Idf(word string) float64 {
	return math.Log(float64(tfidf.numDocs) / (float64(tfidf.wordDocFreq[word]) + 1e-3))
}

func (tfidf *TfIdf) Tf(document []string, word string) float64 {
	var cntWord float64 = 0
	for i := 0; i < len(document); i++ {
		if document[i] == word {
			cntWord++
		}
	}
	return cntWord / float64(len(document))
}

func (tfidf *TfIdf) TfIdf(document []string, word string) float64 {
	return tfidf.Tf(document, word) * tfidf.Idf(word)
}

func (tfidf *TfIdf) TfIdfAllSentence(document []string) []float64 {
	if len(document) == 0 {
		return make([]float64, 0)
	}
	wordCounts := make(map[string]float64)
	for i := 0; i < len(document); i++ {
		wordCounts[document[i]]++
	}
	answer := make([]float64, len(document))
	for i := 0; i < len(document); i++ {
		answer[i] = (wordCounts[document[i]] / float64(len(document))) * tfidf.Idf(document[i])
	}
	return answer
}

package main

import "math"

type TfIdf struct {
	wordDocFreq map[string]int
	numDocs     int
}

func (tfidf *TfIdf) fit(documents [][]string) {
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

func (tfidf *TfIdf) idf(word string) float64 {
	return math.Log(float64(tfidf.numDocs) / float64(tfidf.wordDocFreq[word]))
}

func (tfidf *TfIdf) tf(document []string, word string) float64 {
	var cntWord float64 = 0
	for i := 0; i < len(document); i++ {
		if document[i] == word {
			cntWord++
		}
	}
	return cntWord / float64(len(document))
}

func (tfidf *TfIdf) tfidf(document []string, word string) float64 {
	return tfidf.tf(document, word) * tfidf.idf(word)
}

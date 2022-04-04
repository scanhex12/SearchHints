package tests

import (
	"awesomeProject1/src"
	"math"
	"testing"
)

func TestTfIdf(t *testing.T) {
	tfidf := src.TfIdf{}
	tfidf.Fit([][]string{{"VK", "is", "a", "good", "company"},
		{"I", "am", "using", "VK"},
		{"Snow", "in", "december"}})
	expectedVKIdf := 0.4054651081081644

	const EPS = 1e-9

	if math.Abs(expectedVKIdf-tfidf.Idf("VK")) > EPS {
		t.Error("incorrect IDF implementation")
	}

	expectedVKTfIdf := 0.1013662770270411

	if math.Abs(expectedVKTfIdf-tfidf.TfIdf([]string{"Road", "to", "VK", "intern"}, "VK")) > EPS {
		t.Error("incorrect tf-idf implementation")
	}

}

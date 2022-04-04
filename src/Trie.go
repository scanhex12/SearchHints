package src

type Node struct {
	maxIdf float64
	idWord int
	next   map[byte]*Node
}

type Trie struct {
	root          *Node
	tfidf         *TfIdf
	currentVertex *Node
	allWords      []string
	numWords      int
}

func dfsInsert(v *Node,
	word string,
	iterator int,
	lenght int,
	idfWord float64,
	wordId int) {

	if idfWord > v.maxIdf {
		v.maxIdf = idfWord
		v.idWord = wordId
	}

	if lenght >= iterator {
		return
	}

	if v.next[word[iterator]] == nil {
		v.next[word[iterator]] = &Node{0, 0, make(map[byte]*Node)}
	}
	dfsInsert(v.next[word[iterator]], word, iterator, lenght, idfWord, wordId)
}

func (tree *Trie) InsertWord(word string) {
	idfWord := tree.tfidf.Idf(word)
	tree.numWords++
	tree.allWords = append(tree.allWords, word)
	dfsInsert(tree.root, word, 0, len(word), idfWord, tree.numWords)
}

func (tree *Trie) Fit(documents [][]string) {
	tree.tfidf.Fit(documents)
	for i := 0; i < len(documents); i++ {
		for j := 0; j < len(documents[i]); j++ {
			tree.InsertWord(documents[i][j])
		}
	}
}

func (tree *Trie) ClearInput() {
	tree.currentVertex = tree.root
}

func (tree *Trie) AddSymbol(letter byte) {
	tree.currentVertex = tree.currentVertex.next[letter]
}

func (tree *Trie) CurrentOptimalWord() string {
	return tree.allWords[tree.currentVertex.idWord]
}

func (tree *Trie) OptimalWord(prefixInput string) string {
	currentVertex := tree.root
	for i := 0; i < len(prefixInput); i++ {
		currentVertex = currentVertex.next[prefixInput[i]]
	}
	return tree.CurrentOptimalWord()
}

package on299

type WordDictionary struct {
	children map[rune]*WordDictionary
	isWord   bool
}

func Constructor211() WordDictionary {
	return WordDictionary{
		children: make(map[rune]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	parent := this

	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &WordDictionary{
				children: make(map[rune]*WordDictionary),
			}
			parent.children[ch] = newChild
			parent = newChild
		}
	}

	parent.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	parent := this

	for i, ch := range word {
		if rune(ch) == '.' {
			isMatched := false

			for _, v := range parent.children {
				if v.Search(word[i+1:]) {
					isMatched = true
				}
			}
			return isMatched
		} else if _, ok := parent.children[rune(ch)]; !ok {
			return false
		}

		parent = parent.children[rune(ch)]
	}

	return len(parent.children) == 0 || parent.isWord
}

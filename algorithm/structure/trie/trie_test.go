package trie

import "testing"

func TestTrieInsert(t *testing.T) {
	n := NewNode()

	insettWords := []string{
		"nikola",
		"tesla",
	}

	checkWords := map[string]bool{
		"thomas": false,
		"edison": false,
		"nikola": true,
	}

	n.Insert(insettWords...)
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 2, 12)
}

func TestTrieInsert_substrings(t *testing.T) {
	n := NewNode()

	insertWords := []string{
		"aa",
		"aaaa",
		"aaaaa",
	}

	checkWords := map[string]bool{
		"a":       false,
		"aa":      true,
		"aaa":     false,
		"aaaa":    true,
		"aaaaa":   true,
		"aaaaaa":  false,
		"aaaaaaa": false,
	}

	n.Insert(insertWords...)
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 3, 5+1)

	n.Remove("aaaa")
	checkWords["aaaa"] = false
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 2, 5+1)

	if n.Compact() {
		t.Fatalf("it should not be possible to remove the node")
	}
	n.verify(t, checkWords)
	n.verifySizeCapa(t, 2, 5+1)
}

func TestTrieRemove(t *testing.T) {
	n := NewNode()

	insertWords := []string{
		"nikola",
		"tesla",
		"albert",
		"einstein",
	}

	checkWords := map[string]bool{
		"thomas":   false,
		"edison":   false,
		"nikola":   true,
		"albert":   true,
		"einstein": true,
	}

	n.Insert(insertWords...)
	n.verify(t, checkWords)
	size, capa := n.Size(), n.Capacity()

	n.remove("albert")

	checkWords["albert"] = false
	size--
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	n.Remove("marcel")
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	n.Remove("nikola", "tesla")
	checkWords["nikola"] = false
	checkWords["tesla"] = false
	size -= 2
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	if n.Compact() {
		t.Fatal("the Trie should not be completely removable after compaction")
	}
	if capa <= n.Capacity() {
		t.Fatal("capacity should have reduced following compaction")
	}

	capa = n.Capacity()
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	n.Remove("einstein")
	checkWords["einstein"] = false
	size--
	n.verify(t, checkWords)
	n.verifySizeCapa(t, size, capa)

	if !n.Compact() {
		t.Fatal("the root node of an empty Trie should be marked as removable after compaction")
	}

	n.verifySizeCapa(t, 0, 1)
}

func (n *Node) verify(t *testing.T, checkWords map[string]bool) {

	for k, v := range checkWords {
		ok := n.Find(k)

		if ok != v {
			t.Fatalf(
				"%q is %s supposed to be in the Trie.",
				k,
				map[bool]string{true: "", false: "NOT "}[v],
			)
		}
	}
}

func (n *Node) verifySizeCapa(t *testing.T, expectedSize, expectedCapacity int) {
	if got := n.Size(); got != expectedSize {
		t.Fatalf("Expected Size was %d but got %d", expectedSize, got)
	}
	if got := n.Capacity(); got != expectedCapacity {
		t.Fatalf("Expected Capacity was %d but got %d", expectedCapacity, got)
	}
}

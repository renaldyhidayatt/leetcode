package on399

import (
	"fmt"
	"strconv"
	"unicode"
)

type NestedInteger385 struct {
	Num  int
	List []*NestedInteger385
}

func (n NestedInteger385) IsInteger() bool {
	if n.List == nil {
		return true
	}
	return false
}

func (n NestedInteger385) GetInteger() int {
	return n.Num
}

func (n *NestedInteger385) SetInteger(value int) {
	n.Num = value
}

func (n *NestedInteger385) Add(elem NestedInteger385) {
	n.List = append(n.List, &elem)
}

func (n NestedInteger385) GetList() []*NestedInteger385 {
	return n.List
}

func (n NestedInteger385) Print() {
	if len(n.List) != 0 {
		for _, v := range n.List {
			if len(v.List) != 0 {
				v.Print()
				return
			}
			fmt.Printf("%v ", v.Num)
		}
	} else {
		fmt.Printf("%v ", n.Num)
	}
	fmt.Printf("\n")
}

func isDigital(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func Deserialize(s string) *NestedInteger385 {
	stack, cur := []*NestedInteger385{}, &NestedInteger385{}
	for i := 0; i < len(s); {
		switch {
		case isDigital(s[i]) || s[i] == '-':
			j := 0
			for j = i + 1; j < len(s) && isDigital(s[j]); j++ {
			}
			num, _ := strconv.Atoi(s[i:j])
			next := &NestedInteger385{}
			next.SetInteger(num)
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			} else {
				cur = next
			}
			i = j
		case s[i] == '[':
			next := &NestedInteger385{}
			if len(stack) > 0 {
				stack[len(stack)-1].List = append(stack[len(stack)-1].GetList(), next)
			}
			stack = append(stack, next)
			i++
		case s[i] == ']':
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
		case s[i] == ',':
			i++
		}
	}
	return cur
}

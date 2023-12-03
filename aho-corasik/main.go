package main

import "fmt"

type BorhNode struct {
	childs map[rune]*BorhNode

	parent       *BorhNode
	charToParent rune

	suffixLink          *BorhNode
	efficientSuffixLink *BorhNode

	automateMove map[rune]*BorhNode

	substringIDx int
}

type BorhTrie struct {
	root       *BorhNode
	substrings []string
}

func NewBorhTrie(substrings []string) *BorhTrie {
	root := &BorhNode{
		childs:       map[rune]*BorhNode{},
		automateMove: map[rune]*BorhNode{},
		substringIDx: -1,
	}
	root.suffixLink = root
	root.efficientSuffixLink = root
	trie := &BorhTrie{
		root:       root,
		substrings: substrings,
	}
	for idx, substring := range substrings {
		trie.Insert(idx, substring)
	}
	return trie
}

func (t *BorhTrie) Insert(substringIDx int, s string) {
	node := t.root
	for idx, r := range s {
		sIDx := -1
		if idx == len(s)-1 {
			sIDx = substringIDx
		}

		if _, ok := node.childs[r]; !ok {
			node.childs[r] = &BorhNode{
				childs:       map[rune]*BorhNode{},
				parent:       node,
				charToParent: r,
				automateMove: map[rune]*BorhNode{},
				substringIDx: sIDx,
			}
		}
		node = node.childs[r]
	}
}

func (t *BorhTrie) getSuffixLink(node *BorhNode) *BorhNode {
	if node.suffixLink == nil {
		if node == t.root || node.parent == t.root {
			node.suffixLink = t.root
		} else {
			node.suffixLink = t.automateMove(t.getSuffixLink(node.parent), node.charToParent)
		}
	}
	return node.suffixLink
}

//func (t *BorhTrie) getEfficientSuffixLink(node *BorhNode) *BorhNode {
//	if node.efficientSuffixLink == nil {
//		suffixLink := t.getSuffixLink(node)
//		if suffixLink == t.root || suffixLink.substringIDx != -1 {
//			node.efficientSuffixLink = suffixLink
//		} else {
//			node.efficientSuffixLink = t.getEfficientSuffixLink(node)
//		}
//	}
//	return node.efficientSuffixLink
//}

func (t *BorhTrie) automateMove(node *BorhNode, char rune) *BorhNode {
	if node.automateMove[char] == nil {
		if node.childs[char] != nil {
			node.automateMove[char] = node.childs[char]
		} else if node == t.root {
			node.automateMove[char] = t.root
		} else {
			node.automateMove[char] = t.automateMove(t.getSuffixLink(node), char)
		}
	}
	return node.automateMove[char]
}

func (t *BorhTrie) SearchSubstrings(s string) []string {
	res := make([]string, 0)
	node := t.root
	for idx, char := range s {
		node = t.automateMove(node, char)
		//for tempNode := node; tempNode != t.root; tempNode = t.getEfficientSuffixLink(tempNode) {
		//	if tempNode.substringIDx != -1 {
		//		res = append(res, fmt.Sprintf("%d %s", idx+1-len(t.substrings[tempNode.substringIDx]), t.substrings[tempNode.substringIDx]))
		//	}
		//}
		if node.substringIDx != -1 {
			res = append(res, fmt.Sprintf("%d %s", idx+1-len(t.substrings[node.substringIDx]), t.substrings[node.substringIDx]))
		}
	}
	return res
}

func main() {
	substrings := []string{
		"abc",
		"bcdc",
		"cccb",
		"bcdd",
		"bbbc",
	}

	borhTrie := NewBorhTrie(substrings)

	//s := "helloabhowareyouboboimfineabab!"
	s := "abcdcbcddbbbcccbbbcccbb!"
	for _, res := range borhTrie.SearchSubstrings(s) {
		fmt.Println(res)
	}
}

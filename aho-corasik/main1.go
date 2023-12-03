package main

import (
	"fmt"
)

type BorhNode struct {
	borhNodesByCharacter map[string]*BorhNode
	parent               *BorhNode
	symbolToParent       string
	substringIDx         int // ID of substring "-1" if this node is not the end of substring
	suffixLink           *BorhNode
	efficientSuffixLink  *BorhNode
	autoMove             map[string]*BorhNode
}

type BorhMatrix struct {
	root       *BorhNode
	substrings []string
}

func NewBorhMatrix(substrings []string) *BorhMatrix {
	root := &BorhNode{
		borhNodesByCharacter: map[string]*BorhNode{},
		substringIDx:         -1,
		suffixLink:           nil,
		parent:               nil,
		symbolToParent:       "",
		autoMove:             map[string]*BorhNode{},
	}
	root.suffixLink = root
	return &BorhMatrix{
		substrings: substrings,
		root:       root,
	}
}

func (m *BorhMatrix) Insert(substringIDx int, substring string) {
	currentNode := m.root
	for idx, r := range substring {
		if idx == len(substring)-1 {
			currentNode = insert(currentNode, string(r), substringIDx)
			continue
		}
		currentNode = insert(currentNode, string(r), -1)
	}
}

func insert(node *BorhNode, r string, substringIDx int) *BorhNode {
	res, ok := node.borhNodesByCharacter[r]
	if !ok {
		newNode := &BorhNode{
			borhNodesByCharacter: map[string]*BorhNode{},
			substringIDx:         substringIDx,
			parent:               node,
			symbolToParent:       r,
			autoMove:             map[string]*BorhNode{},
		}
		node.borhNodesByCharacter[r] = newNode
		return newNode
	}
	res.substringIDx = substringIDx
	return res
}

func (m *BorhMatrix) getEfficientSuffixLink(node *BorhNode) *BorhNode {
	if node.efficientSuffixLink == nil {
		suffixLink := m.getSuffixLink(node)
		if suffixLink == m.root || suffixLink.substringIDx != -1 {
			node.efficientSuffixLink = suffixLink
		} else {
			node.efficientSuffixLink = m.getEfficientSuffixLink(suffixLink)
		}
	}
	return node.efficientSuffixLink
}

func (m *BorhMatrix) getSuffixLink(node *BorhNode) *BorhNode {
	if node.suffixLink == nil {
		if node == m.root || node.parent == m.root {
			node.suffixLink = m.root
		} else {
			node.suffixLink = m.getAutoMove(m.getSuffixLink(node.parent), node.symbolToParent)
		}
	}
	return node.suffixLink
}

func (m *BorhMatrix) getAutoMove(node *BorhNode, r string) *BorhNode {
	if node.autoMove[r] == nil {
		if node.borhNodesByCharacter[r] != nil {
			node.autoMove[r] = node.borhNodesByCharacter[r]
		} else if node == m.root {
			node.autoMove[r] = m.root
		} else {
			node.autoMove[r] = m.getAutoMove(m.getSuffixLink(node), r)
		}
	}
	return node.autoMove[r]
}

func (m *BorhMatrix) IsExist(s string) bool {
	currentNode := m.root
	var ok bool
	for _, r := range s {
		if currentNode, ok = currentNode.borhNodesByCharacter[string(r)]; !ok {
			return false
		}
	}
	if currentNode.substringIDx == -1 {
		return false
	}
	return true
}

func (m *BorhMatrix) check(sIDx int, node *BorhNode) []string {
	res := make([]string, 0)
	for ; node != m.root; node = m.getEfficientSuffixLink(node) {
		if node.substringIDx != -1 {
			res = append(res, fmt.Sprintf("%d %s", sIDx+1-len(m.substrings[node.substringIDx]), m.substrings[node.substringIDx]))
		}
	}
	return res
}

func (m *BorhMatrix) FindAllSubstrings(s string) []string {
	res := make([]string, 0)
	currentNode := m.root
	for idx, r := range s {
		currentNode = m.getAutoMove(currentNode, string(r))
		tempNode := currentNode
		res = append(res, m.check(idx, tempNode)...)
	}
	return res
}

/*
abccab

ab
b
cab

a -> b

res = [ab]

*/

func main() {
	subStrings := []string{
		"acac",
		"ac",
		"a",
		"b",
		//"eee",
		//"eeb",
	}

	borhMatrix := NewBorhMatrix(subStrings)
	for idx, substring := range subStrings {
		borhMatrix.Insert(idx, substring)
	}

	//borhMatrix.Print()
	//fmt.Println(borhMatrix.IsExist("abab"))
	for _, s := range borhMatrix.FindAllSubstrings("aclb") {
		fmt.Println(s)
	}
}

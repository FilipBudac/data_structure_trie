package main

import (
	"errors"
	"fmt"
)

const AlphabetSize = 26

type Node struct {
	Name     string
	IsWord   bool
	Children [AlphabetSize]*Node
}

type Tree struct {
	Root *Node
}

func createNode(name string) *Node {
	return &Node{Name: name}
}

func createTree(head *Node) *Tree {
	return &Tree{Root: head}
}

func isWordEmpty(word string) bool {
	return word == ""
}

func getLetterIndex(letter uint8) uint8 {
	return letter - 'A'
}

func (t *Tree) append(word string) (bool, error) {
	if err := errors.New("word is empty"); len(word) <= 0 {
		return false, err
	}

	isWordFound, _ := t.search(word)
	if err := errors.New("word is already in tree"); isWordFound {
		return false, err
	}

	lastNode := insertWord(t.Root, word)
	lastNode.IsWord = true

	return true, nil
}

func (t *Tree) search(word string) (bool, error) {
	if err := errors.New("word is empty"); len(word) <= 0 {
		return false, err
	}

	n := findWord(t.Root, word)
	if n != nil && isWord(n) {
		return true, nil
	}

	return false, nil
}


func (t *Tree) print(words string) {
	printWord(t.Root, words)
}

func (t *Tree) delete(word string) (bool, error) {
	if err := errors.New("word is empty"); len(word) <= 0 {
		return false, err
	}
	isWordInTree, _ := t.search(word)

	if !isWordInTree {
		return false, nil
	}

	deleteWord(t.Root, word)
	isWordInTree, _ = t.search(word)

	return !isWordInTree, nil
}

func deleteWord(n *Node, word string) {
	letterIndex := getLetterIndex(word[0])
	word = word[1:]
	curNode := n.Children[letterIndex]

	if curNode == nil {
		return
	}

	if isWord(curNode) && isWordEmpty(word) && !hasNodeChildren(curNode) {
		n.Children[letterIndex] = nil
	} else if isWord(curNode) && isWordEmpty(word) && hasNodeChildren(n) {
		n.Children[letterIndex].IsWord = false
	} else {
		deleteWord(curNode, word)
	}
}

func printWord(n *Node, words string) {
	if isWord(n) {
		fmt.Println(words)
	}

	for i := 0; i < AlphabetSize; i++ {
		if n.Children[i] != nil {
			printWord(n.Children[i], words + string(rune('A' + i)))
		}
	}
}

func findWord(n *Node, word string) *Node {
	letterIndex := getLetterIndex(word[0])
	curNode := n.Children[letterIndex]

	word = word[1:]

	if isWordEmpty(word) || curNode == nil {
		return curNode
	}

	return findWord(curNode, word)
}

func insertWord(n *Node, word string) *Node {
	letterIndex := getLetterIndex(word[0])

	if n.Children[letterIndex] == nil {
		n.Children[letterIndex] = createNode(string(word[0]))
	}

	word = word[1:]

	if isWordEmpty(word) {
		return n.Children[letterIndex]
	}

	return insertWord(n.Children[letterIndex], word)
}

func hasNodeChildren(n *Node) bool {
	for i := 0; i < AlphabetSize; i++ {
		if n.Children[i] != nil {
			return true
		}
	}
	return false
}

func isWord(n *Node) bool {
	return n.IsWord
}

func main() {
	fmt.Println("Go starting...")

	var head = createNode("")
	var t = createTree(head)

	fmt.Println(t.append("BCCDDX"))
	fmt.Println(t.append("BCCA"))
	fmt.Println(t.append("BCCF"))
	fmt.Println(t.append("CDL"))
	fmt.Println(t.append("BCC"))
	fmt.Println(t.search("OI"))
	fmt.Println(t.delete("BCC"))
}

package morsegode

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// MorseTree describes a binary Morse code tree that can translate between
// sequences of symbols (dots and dashes) and the values they represent
// (letters, digits, punctuation, or any other symbol a code table uses).
//
// A MorseTree has two responsibilities:
//
//   - Decoding: given a sequence of dot/dash symbols, walk the tree to find
//     the value stored at that node (TravelTree) or decode a whole word made
//     up of several Morse letters (DecodeWord).
//   - Encoding: given a value, search the tree for the node that holds it and
//     return the dot/dash path used to reach it (SearchTree) or encode a
//     whole word (EncodeWord).
type MorseTree interface {
	// DecodeWord decodes a single word made up of Morse letters, where each
	// element of LetterSeq is one Morse letter (for example "...---...").
	DecodeWord(LetterSeq []string) (string, error)

	// TravelTree walks the tree following the given dot/dash symbols and
	// returns the value stored at the resulting node.
	TravelTree(MorseLetter []byte) (string, error)

	// EncodeWord encodes a word made up of values (for example letters),
	// joining the encoded letters with LetterJoin.
	EncodeWord(Letters []string, LetterJoin string) (string, error)

	// SearchTree finds the value Letter in the tree and returns the dot/dash
	// path used to reach it.
	SearchTree(Letter string) (string, error)
}

// GeneralMorseTree is a binary tree implementation of the MorseTree
// interface. Each level of the tree corresponds to one symbol: the left child
// is reached with LeftSymbol and the right child with RightSymbol. The value
// stored at the root is the empty string.
type GeneralMorseTree struct {
	// TreeRoot is the root node of the tree.
	TreeRoot MorseTreeNode

	// LeftSymbol is the symbol that moves down the left branch of the tree
	// (typically '.').
	LeftSymbol byte

	// RightSymbol is the symbol that moves down the right branch of the tree
	// (typically '-').
	RightSymbol byte
}

// NewGeneralMorseTree builds a GeneralMorseTree from a code table. The table
// maps a value (for example the letter "s") to the sequence of symbols that
// encodes it (for example "..."). The tree stores the value at the node
// reached by following each symbol of the sequence.
//
// left and right are the symbols used to descend the left and right branches
// of the tree, typically '.' and '-'.
//
// Any value whose code contains a symbol that is neither left nor right
// results in an error, as does an empty code.
func NewGeneralMorseTree(left, right byte, table map[string]string) (*GeneralMorseTree, error) {
	root := &GeneralMorseTreeNode{}

	for value, path := range table {
		if path == "" {
			return nil, fmt.Errorf("morse path for value %q cannot be empty", value)
		}

		node := root
		for i := 0; i < len(path); i++ {
			symbol := path[i]
			switch symbol {
			case left:
				if node.left == nil {
					node.left = &GeneralMorseTreeNode{}
				}
				node = node.left
			case right:
				if node.right == nil {
					node.right = &GeneralMorseTreeNode{}
				}
				node = node.right
			default:
				return nil, fmt.Errorf(
					"invalid symbol %q in path %q (want %q or %q)",
					symbol, path, string(left), string(right),
				)
			}
		}

		node.letter = value
	}

	return &GeneralMorseTree{
		TreeRoot:    root,
		LeftSymbol:  left,
		RightSymbol: right,
	}, nil
}

// DecodeWord decodes a single word made up of Morse letters. Each element of
// LetterSeq is one Morse letter (for example "...---..."), and the decoded
// values are joined into a single string without any separator.
func (tree GeneralMorseTree) DecodeWord(LetterSeq []string) (string, error) {
	word := make([]string, 0)
	for _, morseSeq := range LetterSeq {
		if morseSeq == "" {
			continue
		}

		letter, err := tree.TravelTree([]byte(morseSeq))
		if err != nil {
			return "", err
		}
		word = append(word, letter)
	}

	return strings.Join(word, ""), nil
}

// TravelTree walks the tree following the given dot/dash symbols and returns
// the value stored at the resulting node. An error is returned if the
// sequence contains a symbol that is neither the tree's LeftSymbol nor its
// RightSymbol, or if the sequence leads to a node that does not exist in the
// tree.
func (tree GeneralMorseTree) TravelTree(MorseLetter []byte) (string, error) {
	currentNode := tree.TreeRoot
	if isNilNode(currentNode) {
		return "", errors.New("invalid sequence of symbols")
	}

	for _, symbol := range MorseLetter {
		switch symbol {
		case tree.LeftSymbol:
			currentNode = currentNode.Left()
		case tree.RightSymbol:
			currentNode = currentNode.Right()
		default:
			return "", fmt.Errorf("invalid symbol %q in morse sequence", symbol)
		}

		if isNilNode(currentNode) {
			return "", errors.New("invalid sequence of symbols")
		}
	}

	value, ok := currentNode.Item().(string)
	if !ok {
		return "", errors.New("node contains no value")
	}
	return value, nil
}

// isNilNode reports whether node is nil, including the case where an
// interface holds a typed nil pointer (for example a nil
// *GeneralMorseTreeNode stored in a MorseTreeNode interface).
func isNilNode(node MorseTreeNode) bool {
	if node == nil {
		return true
	}

	v := reflect.ValueOf(node)
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return v.IsNil()
	}
	return false
}

// EncodeWord encodes a word made up of values (for example letters), joining
// the encoded letters with LetterJoin. Empty values are skipped.
func (tree GeneralMorseTree) EncodeWord(Letters []string, LetterJoin string) (string, error) {
	word := make([]string, 0)

	for _, letter := range Letters {
		if letter == "" {
			continue
		}

		encodedLetter, err := tree.SearchTree(letter)
		if err != nil {
			return "", err
		}

		word = append(word, encodedLetter)
	}

	return strings.Join(word, LetterJoin), nil
}

// SearchTree finds the value Letter in the tree and returns the dot/dash path
// used to reach it. Searching for the empty string returns an error.
func (tree GeneralMorseTree) SearchTree(letter string) (string, error) {
	if letter == "" {
		return "", errors.New("letter not found")
	}

	node, ok := tree.TreeRoot.(*GeneralMorseTreeNode)
	if !ok {
		return "", errors.New("unsupported tree node type")
	}

	return tree.searchNode(node, "", letter)
}

// searchNode searches recursively for letter, building up the symbol path as
// it descends. It returns the first node in a depth-first, left-to-right order
// whose value matches letter.
func (tree GeneralMorseTree) searchNode(
	node *GeneralMorseTreeNode,
	path string,
	letter string,
) (string, error) {
	if node == nil {
		return "", errors.New("letter not found")
	}

	if node.letter == letter {
		return path, nil
	}

	// Search left
	if result, err := tree.searchNode(
		node.left,
		path+string(tree.LeftSymbol),
		letter,
	); err == nil {
		return result, nil
	}

	// Search right
	if result, err := tree.searchNode(
		node.right,
		path+string(tree.RightSymbol),
		letter,
	); err == nil {
		return result, nil
	}

	return "", errors.New("letter not found")
}

// MorseTreeNode is a single node in a Morse code tree. Its two children are
// reached with the tree's left and right symbols, and its Item is the value
// stored at this node.
type MorseTreeNode interface {
	// Left returns the left child node.
	Left() MorseTreeNode

	// Right returns the right child node.
	Right() MorseTreeNode

	// Item returns the value stored at this node.
	Item() any
}

// GeneralMorseTreeNode is the default MorseTreeNode implementation. It stores
// a value as a string and has two optional children, left and right.
type GeneralMorseTreeNode struct {
	// letter is the value stored at this node.
	letter string

	left  *GeneralMorseTreeNode
	right *GeneralMorseTreeNode
}

// Left returns the left child node, or nil if there is none.
func (node *GeneralMorseTreeNode) Left() MorseTreeNode {
	if node == nil || node.left == nil {
		return nil
	}
	return node.left
}

// Right returns the right child node, or nil if there is none.
func (node *GeneralMorseTreeNode) Right() MorseTreeNode {
	if node == nil || node.right == nil {
		return nil
	}
	return node.right
}

// Item returns the value stored at this node.
func (node *GeneralMorseTreeNode) Item() any {
	return node.letter
}

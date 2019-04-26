package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	childCount    int
	metadataCount int
	children      []*node
	metadata      []int
	parent        *node
}

func (n *node) sum() int {
	if len(n.children) == 0 {
		sum := 0
		for _, value := range n.metadata {
			sum += value
		}
		return sum
	} else {
		sum := 0
		for _, value := range n.metadata {
			i := value - 1
			if i < 0 || i >= len(n.children) {
				continue
			}
			sum += n.children[i].sum()
		}
		return sum
	}
}

func pop(in []int) ([]int, int, bool) {
	if len(in) == 0 {
		return in, 0, false
	}
	return in[1:], in[0], true
}

var nodes = []*node{}
var parent *node
var nums = readFile()

func main() {
	root, ok := createNode(parent)
	_ = ok
	_ = root
	fmt.Println(root.sum())
}

func createNode(parent *node) (*node, bool) {
	node, ok := createNodeWithHeader(parent)
	if !ok {
		return nil, false
	}
	for node.childCount != len(node.children) {
		child, ok := createNode(node)
		if !ok {
			return nil, false
		}
		node.children = append(node.children, child)
	}
	for node.metadataCount != len(node.metadata) {
		var current int
		var ok bool

		nums, current, ok = pop(nums)
		if !ok {
			return nil, false
		}
		node.metadata = append(node.metadata, current)
	}
	return node, true
}

func createNodeWithHeader(parent *node) (*node, bool) {
	var current int
	var ok bool

	nums, current, ok = pop(nums)
	if !ok {
		return nil, false
	}
	n := &node{}
	n.children = []*node{}
	n.metadata = []int{}
	n.parent = parent
	nodes = append(nodes, n)
	parent = n
	n.childCount = current
	nums, current, ok = pop(nums)
	if !ok {
		return nil, false
	}
	n.metadataCount = current
	return n, true
}

func readFile() []int {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	texts := strings.Split(string(b), " ")
	numbers := make([]int, len(texts))
	for i, text := range texts {
		num, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}
	return numbers
}

package main

import (
	"fmt"

	"github.com/iamseki/non-linear-ds/binary_tree"
)

func main() {
	// nodes with no more than 2 children
	// binary search tree:
	//
	//				parent
	//			/				\
	//	smaller		larger
	//
	// well balanced tree has O(log n) time complexity for search, delete and insert operations
	tree := &binary_tree.Node{Value: 100}
	tree.Insert(120)
	tree.Insert(10)
	tree.Insert(17)
	tree.Insert(11)
	tree.Insert(12)

	fmt.Printf("search for %v, result: %v\n", 10, tree.Search(10))
	fmt.Printf("search for %v, result: %v\n", 17, tree.Search(17))
	fmt.Printf("search for %v, result: %v\n", 11, tree.Search(1100))
}

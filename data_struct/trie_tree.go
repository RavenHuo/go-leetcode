/**
 * @Author raven
 * @Description
 * @Date 2022/10/31
 **/
package data_struct

import "sort"

type TrieTree struct {
	path     string
	children []*TrieTree
	value    interface{}
}

func (tree TrieTree) Add(path string, value interface{}) {
	sort.Slice(tree.children, func(i, j int) bool {
		return false
	})
}

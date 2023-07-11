package gee

import (
	"strings"
)

type node struct {
	pattern  string //待匹配路由 /p/:lang
	part     string //路由中的一部分 :lang
	children []*node
	isWild   bool //是否宽范围的匹配 // : *可以
}

func (n *node) matchNode(part string) *node {

	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}

	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]

	child := n.matchNode(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasSuffix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]

	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result

		}
	}
	return nil
}

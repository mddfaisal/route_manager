// @author Mohd Faisal
// @license BSD-3-Clause
// copyright 2025 Mohd Faisal
// @version 1.0.0
// @description This package provides a way to manage routes in a tree structure.

package routemanager

import (
	"fmt"
	"regexp"
	"strings"
)

var current, prev *Node

type Node struct {
	Resource  interface{}
	Element   string
	IsDynamic bool
	Node      []Node
}

func splitAndValidate(s string) []string {
	result := strings.Split(s, "/")
	if len(result[0]) != 0 {
		panic("Invalid path: ")
	}
	result = result[1:] //
	if result[0] == ":" {
		panic("Invalid path: " + s + " cannot start with a dynamic segment")
	}
	if len(result) == 1 && result[0] == "" {
		panic("Invalid path: " + s + " cannot be empty")
	}
	if len(result) > 1 && result[0] == "" {
		panic("Invalid path: " + s + " cannot have multiple segments without a leading slash")
	}
	for i := range result {
		if len(result[i]) == 0 {
			panic("Invalid path: " + s + " at index " + fmt.Sprintf("%v", i))
		}
		if !regexp.MustCompile("^[:a-zA-Z0-9_-]*$").MatchString(result[i]) {
			panic("Invalid path: " + s + " at index " + fmt.Sprintf("%v", i))
		}
	}
	return result
}

func search(r *Node, exp []string) string {
	match := ""
	if len(exp) == 1 && r.Element == exp[0] {
		return r.Element
	}
	if r.Element == string(exp[0]) {
		exp = exp[1:]
		for _, n := range r.Node {
			if n.IsDynamic && strings.HasPrefix(exp[0], ":") {
				panic("Invalid path: " + strings.Join(exp, "/") + " cannot start with a dynamic segment")
			}
			if n.Element == string(exp[0]) {
				return search(&n, exp)
			}
		}
	}
	return match
}

// Search traverses the tree to find a node that matches the given expression.
// It returns the element if found, otherwise an empty string.
func Search(r *Node, exp string) string {
	splt := splitAndValidate(exp)
	matched := ""
	if len(splt) == 1 && r.Element == splt[0] {
		return r.Element
	}
	if len(r.Node) > 0 {
		for _, node := range r.Node {
			if node.IsDynamic && strings.HasPrefix(splt[0], ":") {
				panic("Invalid path: " + exp + " cannot start with a dynamic segment")
			}
			if node.Element == string(splt[0]) {
				matched = search(&node, splt)
				if len(matched) > 0 {
					break
				}
			}
		}
	}
	return matched
}

// setNode sets the current node based on the element and index.
// If the index is 0, it sets the node to the first matching element.
// If the index is greater than 0, it traverses the tree to find the node.
// If strict is true, it only matches the exact element.
func setNode(r *Node, node **Node, exp string, strict bool, index int) {
	if index == 0 {
		for i := range r.Node {
			if r.Node[i].Element == exp {
				*node = &r.Node[i]
				return
			}
		}
		*node = r
		return
	}
	// If strict is true, we only match the exact element.
	// If strict is false, we allow matching the first element or an empty node.
	if strict {
		if r.Element == exp {
			*node = r
			return
		}
	} else {
		if r.Element == exp || len(r.Node) == 0 {
			if r.IsDynamic && strings.HasPrefix(exp, ":") {
				panic("Invalid path: " + exp + " cannot start with a dynamic segment")
			}
			*node = r
			return
		}
	}
	for i := range r.Node {
		setNode(&r.Node[i], node, exp, strict, index)
	}
}

// Insert adds a new path to the tree structure.
// It validates the path, splits it into segments, and inserts it into the tree.
// It also validates dynamic segments and ensures the path structure is correct.
// It panics if the path is invalid or if a dynamic segment is at the start of the path.
// It uses a recursive approach to traverse the tree and insert the new node.
// It also handles the case where the path is empty or has multiple segments without a leading slash.
// It panics if the path is invalid or if a dynamic segment is at the start of the path.
// It uses a recursive approach to traverse the tree and insert the new node.
// It also handles the case where the path is empty or has multiple segments without a leading slash.
func Insert(r *Node, exp string) {
	splt := splitAndValidate(exp)
	for i := range splt {
		node := Node{
			Element:   splt[i],
			IsDynamic: strings.HasPrefix(splt[i], ":"),
			Node:      []Node{},
		}
		current = nil
		setNode(r, &current, splt[i], false, i)
		if current.Element == splt[i] {
			continue
		}
		if i > 0 {
			prev = nil
			setNode(r, &prev, splt[i-1], true, i)
			if prev.Element == splt[i-1] {
				prev.Node = append(prev.Node, node)
			}
			continue
		}
		current.Node = append(current.Node, node)
	}
}

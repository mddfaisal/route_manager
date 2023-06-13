package main

import "fmt"

type Node struct {
	Element string
	Node    []Node
}

func CreateTree(root *Node) {
	root.Element = "/"
	root.Node = []Node{}

	for _, e := range []string{"A", "B"} {
		root.Node = append(root.Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}

	// insert B, C, D into A & M into B
	root.Node[1].Node = append(root.Node[1].Node, Node{Element: "M", Node: []Node{}})
	root.Node[1].Node[0].Node = append(root.Node[1].Node[0].Node, Node{
		Element: "N",
		Node:    []Node{},
	})

	for _, e := range []string{"B", "C", "D"} {
		root.Node[0].Node = append(root.Node[0].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}

	for _, e := range []string{"F", "G"} {
		root.Node[0].Node[0].Node = append(root.Node[0].Node[0].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}

	for _, e := range []string{"O", "P", "Q", "R"} {
		root.Node[0].Node[0].Node[1].Node = append(root.Node[0].Node[0].Node[1].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}
	root.Node[0].Node[0].Node[0].Node = append(root.Node[0].Node[0].Node[0].Node, Node{
		Element: "N",
		Node:    []Node{},
	})

	for _, e := range []string{"H", "I", "J"} {
		root.Node[0].Node[1].Node = append(root.Node[0].Node[1].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}
	root.Node[0].Node[1].Node[0].Node = append(root.Node[0].Node[1].Node[0].Node, Node{
		Element: "S",
		Node:    []Node{},
	})

	root.Node[0].Node[1].Node[1].Node = append(root.Node[0].Node[1].Node[1].Node, Node{
		Element: "T",
		Node:    []Node{},
	})
	for _, e := range []string{"K", "M", "L"} {
		root.Node[0].Node[2].Node = append(root.Node[0].Node[2].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}

	for _, e := range []string{"U", "V"} {
		root.Node[0].Node[2].Node[0].Node = append(root.Node[0].Node[2].Node[0].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}
	for _, e := range []string{"X", "Y"} {
		root.Node[0].Node[2].Node[1].Node = append(root.Node[0].Node[2].Node[1].Node, Node{
			Element: e,
			Node:    []Node{},
		})
	}
	root.Node[0].Node[2].Node[1].Node[1].Node = append(root.Node[0].Node[2].Node[1].Node[0].Node, Node{
		Element: "Z",
		Node:    []Node{},
	})
	root.Node[0].Node[2].Node[2].Node = append(root.Node[0].Node[2].Node[2].Node, Node{
		Element: "W",
		Node:    []Node{},
	})
}

func Display(n *Node) {
	fmt.Println(n.Element)
	if len(n.Node) > 0 {
		for _, node := range n.Node {
			Display(&node)
		}
	}
}

func main() {
	root := new(Node)
	CreateTree(root)
	Display(root)
	fmt.Println(root)
}

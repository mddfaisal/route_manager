package utils

var Exp = []string{
	"ABFN",
	"ABGO",
	"ABGP",
	"ABGR",
	"ABGQ",
	"ACHS",
	"ACIT",
	"ACJ",
	"ADKU",
	"ADKV",
	"ADLW",
	"ADMX",
	"ADMYZ",
	"BMN",
}

type Node struct {
	Element string
	Node    []Node
}

// var node *Node

func Split(s string) []string {
	e := []string{}
	for _, s := range s {
		e = append(e, string(s))
	}
	return e
}

func search(r *Node, exp string) string {
	match := ""
	if len(exp) == 1 && r.Element == exp {
		return r.Element
	}
	if r.Element == string(exp[0]) {
		exp = exp[1:]
		for _, n := range r.Node {
			if n.Element == string(exp[0]) {
				return search(&n, exp)
			}
		}
	}
	return match
}

func Search(r *Node, exp string) string {
	matched := ""
	if len(exp) == 1 && r.Element == exp {
		return r.Element
	}
	if len(r.Node) > 0 {
		for _, node := range r.Node {
			if node.Element == string(exp[0]) {
				matched = search(&node, exp)
				if len(matched) > 0 {
					break
				}
			}
		}
	}
	return matched
}

// func insert(r *Node, exp string) {
// 	if len(exp) == 0 || r == nil {
// 		return
// 	}
// 	// root element case
// 	if r.Element == "/" && len(r.Node) == 0 {
// 		r.Node = append(r.Node, Node{Element: string(exp[0]), Node: []Node{}})
// 		insert(&r.Node[0], exp[1:])
// 		return
// 	}
// 	if len(r.Node) > 0 {
// 		for i := range r.Node {
// 			if r.Node[i].Element == string(exp[0]) {
// 				exp = exp[1:]
// 			}
// 			insert(&r.Node[i], exp)
// 		}
// 	} else {
// 		r.Node = append(r.Node, Node{Element: string(exp[0]), Node: []Node{}})
// 		insert(r, exp[1:])
// 		return
// 	}

// }

// func Insert(r *Node, exp string) {
// 	splt := Split(exp)
// 	for i := range splt {
// 		fmt.Println(splt[i])
// 	}
// }

// func search2(r *Node, exp string) {
// 	if r.Element == exp {
// 		node = r
// 	}
// 	for i := range r.Node {
// 		search2(&r.Node[i], exp)
// 	}
// }

// func Search2(r *Node, exp string) {
// 	fmt.Println("searching exp: ", exp)
// 	for i := range exp {
// 		search2(r, string(exp[i]))
// 		if node.Element == string(exp[i]) {
// 			fmt.Println("route found: ")
// 		}
// 	}
// }

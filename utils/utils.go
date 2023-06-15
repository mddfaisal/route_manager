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
		match = r.Element
		return match
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

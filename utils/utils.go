package utils

var (
	Exp = []string{
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
	current, prev *Node
)

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

func setNode(r *Node, node **Node, exp string, strict bool) {
	if strict {
		if r.Element == exp {
			*node = r
			return
		}
	} else {
		if r.Element == exp || len(r.Node) == 0 {
			*node = r
			return
		}
	}
	for i := range r.Node {
		setNode(&r.Node[i], node, exp, strict)
	}
}

func Insert(r *Node, exp string) {
	splt := Split(exp)
	for i := range splt {
		node := Node{Element: splt[i], Node: []Node{}}
		current = nil
		setNode(r, &current, splt[i], false)
		if current.Element == splt[i] {
			continue
		}
		if i > 0 {
			prev = nil
			setNode(r, &prev, splt[i-1], true)
			if prev.Element == splt[i-1] {
				prev.Node = append(prev.Node, node)
			}
			continue
		}
		current.Node = append(current.Node, node)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"tree/utils"
)

func printRoot(root *utils.Node) {
	b, err := json.Marshal(*root)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {

	r1 := &utils.Node{Element: "/", Node: []utils.Node{}}
	for i := range utils.Exp {
		utils.Insert(r1, utils.Exp[i])
	}
	printRoot(r1)
	for i := range utils.Exp {
		fmt.Println(utils.Exp[i], " ===> ", utils.Search(r1, utils.Exp[i]))
	}
}

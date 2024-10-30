package main

import (
	"fmt"

	"github.com/jerryzhen01/goeapi"
)

func main() {
	node, err := goeapi.ConnectTo("dut")
	if err != nil {
		panic(err)
	}
	commands := []string{"show version"}
	conf, err := node.Enable(commands)
	if err != nil {
		panic(err)
	}
	for k, v := range conf[0] {
		fmt.Println("k:", k, "v:", v)
	}
}

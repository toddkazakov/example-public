package main

import (
    "fmt"
    _ "github.com/toddkazakov/example-public/public"
    "github.com/toddkazakov/example-public/registry"
)

func main() {
	fmt.Println("Registered modules: ")
	for _, m := range registry.GetModules() {
		fmt.Println(m)
	}
}

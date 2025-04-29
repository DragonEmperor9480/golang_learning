package main

import "fmt"

type Engine struct {
	horsePower int
}

type Car struct {
	brand  string
	model  string
	engine Engine
}

func main()  {
	
	c1:= Car{
		brand: "Ford",
		model: "Mustang",
		engine: Engine{
			horsePower: 450,
		},
	}

	fmt.Println(c1.brand,c1.model,c1.engine.horsePower)
}
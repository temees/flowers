package main

import (
	"fmt"
	"les3/flowers"
)

func main() {
	a := []flowers.Flower{
		flowers.Flower{
			Id:    1,
			Name:  "rose",
			Price: 30,
		},
		flowers.Flower{
			Id:    2,
			Name:  "tulip",
			Price: 20,
		},
	}
	fmt.Println(a)
}

package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"les3/flowers"
	"log"
)

type Config struct {
	Debug bool
	Port  int
}

func main() {
	var s Config
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\n"
	_, err = fmt.Printf(format, s.Debug, s.Port)
	if err != nil {
		log.Fatal(err.Error())
	}
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

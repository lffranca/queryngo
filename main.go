package main

import (
	"github.com/lffranca/queryngo/pkg/queryingo"
	"log"
)

func main() {
	if err := queryingo.Run(); err != nil {
		log.Panicln(err)
	}
}

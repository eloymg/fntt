package main

import (

	"github.com/eloymg/fntt/reader"
	"os"

)

func main() {
	path := os.Args[1]
	r := reader.NewReader()
	r.Read(path)
}

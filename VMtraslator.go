package main

import (

	"local/VMtranslator/reader"
	"os"

)

func main() {

	path := os.Args[1]
	r := reader.NewReader()
	r.Read(path)
}

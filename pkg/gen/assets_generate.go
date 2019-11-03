// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	fs := http.Dir("./templates")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "gen",
		VariableName: "templates",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

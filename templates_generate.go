// +build ignore

package main

import (
	"github.com/sdorra/jasas/templates"
	"github.com/shurcooL/vfsgen"
	"log"
)

func main() {
	err := vfsgen.Generate(templates.Templates, vfsgen.Options{
		PackageName:  "templates",
		BuildTags:    "!dev",
		VariableName: "Templates",
		Filename:     "templates/templates_prod.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

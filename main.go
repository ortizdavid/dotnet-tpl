package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ortizdavid/dotnet-tpl/generators"
	"github.com/ortizdavid/go-nopain/messages"
)

func main() {
	templateType := flag.String("template", "", "Template type (mvc/webapi)")
	flag.Parse()

	if *templateType == "" {
		fmt.Println("Usage: dotnet-tpl --template <TEMPLATE> --name <NAME>")
		os.Exit(1)
	}

	var generator generators.Generator
	fmt.Printf("Generating template '%s'\n", *templateType)
	if err := generator.GenerateTemplate(*templateType); err != nil {
		messages.Error(err.Error())
		os.Exit(1)
	}
	messages.Success(fmt.Sprintf("'%s' tempalte created successfully!", *templateType))
}
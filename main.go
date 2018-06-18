package main

import (
	"flag"
	"fmt"

	"github.com/mylesnoton/password-generator/generator"
)

func main() {
	length := flag.Int("length", 16, "The length of password you want")
	noSpecial := flag.Bool("no-special", false, "ff")

	flag.Parse()

	pwd, err := generator.NewPassword(*length, *noSpecial)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(pwd)
}

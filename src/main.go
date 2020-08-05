package main

import (
	"flag"
	"fmt"
)

func main() {
	help := flag.Bool("h", false, "quest --h")
	flag.Parse()

	if *help {
		fmt.Println("Quest! The Developer To-Do App")
		fmt.Println()
		fmt.Println("USAGE:")
		fmt.Println("\tquest [COMMAND] [SUBCOMMAND]")
		fmt.Println("COMMANDS:")
		fmt.Println("\tnew")
		fmt.Println("\tlist")
		fmt.Println("\thelp")
	} else {
		fmt.Println("Welcome to the Quest!")
	}
}

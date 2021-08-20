package main

import (
	"flag"

	"github.com/mantissts/InfoGopher/modules/filesystem"
	"github.com/mantissts/InfoGopher/modules/memory"
)

func main() {

	directory := flag.String("directory", "", "Specify the root directory for the installed application")
	memorySearchString := flag.String("mem-search", "", "A string to search in the memory")
	flag.Parse()

	filesystem.GetFiles(*directory)
	memory.FindStringInMemory(*memorySearchString)
}

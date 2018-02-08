package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hinshun/libpackage"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) == 1 {
		return fmt.Errorf("no args")
	}

	_, reverse, err := libpackage.ComputeGraph()
	if err != nil {
		return err
	}

	reachable := reverse.Search(os.Args[1:]...)
	for pkg := range reachable {
		fmt.Printf("%s\n", pkg)
	}

	return nil
}

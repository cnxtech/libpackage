package libpackage

import (
	"fmt"
	"go/build"

	"golang.org/x/tools/refactor/importgraph"
)

func ComputeGraph() (forward, reverse importgraph.Graph, err error) {
	var errs map[string]error
	forward, reverse, errs = importgraph.Build(&build.Default)
	if errs != nil {
		for path, err := range errs {
			fmt.Printf("[%s]: %s\n", path, err)
		}
		return nil, nil, fmt.Errorf("err")
	}

	return forward, reverse, nil
}

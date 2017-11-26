package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/keitax/tv-cptl"
	"github.com/pkg/errors"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", errors.Wrap(err, "failed to open stdin"))
		os.Exit(1)
	}
	fmt.Print(tvcptl.RenderBlocks(tvcptl.ParseBlocks(string(bs))))
}

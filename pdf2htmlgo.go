package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
)

var opts struct {
	Version []bool `short:"v" long:"version" description:"print copyright and version info"`
}

func main() {
	// Parse flags
	args, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Version: %d\n", len(opts.Version))
	fmt.Printf("Remaining args: %s\n", strings.Join(args, " "))
}

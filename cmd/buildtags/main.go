package main

import (
	"fmt"

	"github.com/jomoespe/webc-spikes/pkg/buildtags"
)

func main() {
	fmt.Printf("Hello %s!\n", buildtags.Salutation())
}

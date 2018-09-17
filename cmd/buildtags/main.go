// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/jomoespe/webc-spikes/pkg/buildtags"
)

func main() {
	fmt.Printf("Hello %s!\n", buildtags.Salutation())
}

// Copyright 2010 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

package composer_test

import (
	"fmt"
	"os"

	"github.com/jomoespe/webc-spikes/pkg/composer"
)

func ExampleCompose() {
	filename := "../../testdata/process-html/example-1.html"

	b, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot open html test file %s", filename)
	}
	defer b.Close()
	doc, err := composer.Parse(b)
	if err != nil {
		fmt.Printf("cannot parse html test file %s", filename)
	}

	doc.Compose()

	// Output:
	// section
	// section
	// section
}

// Copyright 2018 Jomoespe. All rights reserved.
// Use of this source code is governed by a WTFPL-style
// license that can be found in the LICENSE file.

// +build tag1 !tag2

package buildtags

func Salutation() (name string) {
	name = "from tag1"
	return
}

// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	_ "github.com/2637309949/bulrush-cmd/command"
	"github.com/devfacet/gocmd"
)

func main() {
	gocmd.New(gocmd.Options{
		Name:        "bul",
		Version:     "1.0.0",
		Description: "bulrush application generator tool",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}

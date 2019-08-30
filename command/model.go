// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"github.com/devfacet/gocmd"
)

// Model command
var _, _ = gocmd.HandleFlag("New.Model", func(cmd *gocmd.Cmd, args []string) error {
	// fmt.Println(math.Pow(flags.New.Pow.Base, flags.New.Pow.Exponent))
	return nil
})

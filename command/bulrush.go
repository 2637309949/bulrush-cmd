// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package command

import (
	"fmt"

	"github.com/devfacet/gocmd"
)

// Echo command
var _, _ = gocmd.HandleFlag("Init", func(cmd *gocmd.Cmd, args []string) error {
	fmt.Printf("%s\n", cmd.FlagArgs("Init")[len(cmd.FlagArgs("Init"))-1])
	return nil
})

// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

// Flags defined all flags to use
type Flags struct {
	Help      bool `short:"h" long:"help" description:"Display usage" global:"true"`
	Version   bool `short:"v" long:"version" description:"Display version"`
	VersionEx bool `long:"vv" description:"Display version (extended)"`
	Init      struct {
		Plugin float64 `short:"p" long:"plugin" required:"false" description:"Use the plug-in"`
		Dir    string  `short:"d" long:"dir" required:"true" description:"Dir for generator"`
	} `command:"init" description:"Create a project"`
	New struct {
		Plugin struct {
			Dir  string `short:"d" long:"dir" required:"false" description:"Dir for generator"`
			Name string `short:"n" long:"name" required:"true" description:"Name for generator"`
		} `command:"plugin" description:"Create plugins"`
		Route struct {
			Dir  string `short:"d" long:"dir" required:"false" description:"Dir for generator"`
			Name string `short:"n" long:"name" required:"true" description:"Name for generator"`
		} `command:"route" description:"Create routes"`
		Model struct {
			Dir  string `short:"d" long:"dir" required:"false" description:"Dir for generator"`
			Name string `short:"n" long:"name" required:"true" description:"Name for generator"`
		} `command:"model" description:"Create models"`
	} `command:"new" description:"New routes, models or plugins" nonempty:"true"`
}

var flags = Flags{}

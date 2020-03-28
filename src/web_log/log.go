// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package web_log

import (
	"github.com/orivil/novel-spider/src/env"
	"io"
	"log"
	"os"
)

type flag int

const (
	FlagInfo flag = 1 << iota
	FlagError
	FlagWarning
)

var Info = log.New(os.Stdout, "[info] ", 0)

var Warning = log.New(os.Stderr, "[warning] ", 0)

var Error = log.New(os.Stderr, "[error] ", 0)

func init() {
	if env.CONFIG.DEBUG {
	}
}

func SetOutput(flag flag, writer io.Writer) {
	if flag&FlagInfo != 0 {
		Info.SetOutput(writer)
	}
	if flag&FlagWarning != 0 {
		Warning.SetOutput(writer)
	}
	if flag&FlagError != 0 {
		Error.SetOutput(writer)
	}
}

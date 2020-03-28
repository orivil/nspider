// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package websocket

import (
	"github.com/orivil/novel-spider/src/web_log"
	"io"
	"sync"
)

var Console = NewConsole()

func init() {
	web_log.SetOutput(web_log.FlagInfo | web_log.FlagError | web_log.FlagWarning, Console)
}

type console struct {
	connects map[string]io.Writer
	mu sync.Mutex
}

func NewConsole() *console {
	return &console{
		connects: make(map[string]io.Writer, 1),
		mu:       sync.Mutex{},
	}
}

func (c *console) AddClient(key string, client io.Writer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.connects[key] = client
}

func (c *console) Write(p []byte) (n int, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, conn := range c.connects {
		conn.Write(p)
	}
	return 0, nil
}

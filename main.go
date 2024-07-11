package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	g "xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/shockwave/out"
	"xabbo.b7c.io/goearth/shockwave/room"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "Talk To Shout",
	Description: "Shout To Talk, reverses the functionality for ease of use :)",
	Version:     "1.1.1",
	Author:      "Eduard, b7",
})

var roomMgr = room.NewManager(ext)

func main() {
	ext.Intercept(out.CHAT).With(handleTalk)
	ext.Intercept(out.SHOUT).With(handleShout)
	ext.Run()
}

func handleShout(e *g.Intercept) {
	e.Packet.Header = ext.Headers().Get(out.CHAT)
	if handleCommands(e.Packet.ReadString()) {
		e.Block()
	}
}
func handleTalk(e *g.Intercept) {
	e.Packet.Header = ext.Headers().Get(out.SHOUT)
	test := handleCommands(e.Packet.ReadString())
	fmt.Println(test)
	if test {
		e.Block()
	}
}

func handleCommands(msg string) bool {
	if msg == ":dance" {
		ext.Send(out.DANCE)
		return true
	} else if msg == ":pickall" {
		if roomMgr.IsOwner {
			go func() {
				for _, obj := range roomMgr.Objects {
					ext.Send(out.ADDSTRIPITEM, []byte("new stuff "+strconv.Itoa(obj.Id)))
					time.Sleep(time.Millisecond * 500)
				}
			}()
			return true
		}
		return false
	} else if strings.Contains(msg, "o/") {
		ext.Send(out.WAVE)
		return msg == "o/"
	}
	return false
}

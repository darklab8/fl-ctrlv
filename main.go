package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	_ "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"golang.design/x/clipboard"
)

func main() {
	add()
}

func add() {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"v", "ctrl"}, func(e hook.Event) {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("ctrl-v IS PRESSED")

		text := clipboard.Read(clipboard.FmtText)
		// fmt.Println(string(text))
		robotgo.TypeStr(string(text))
		fmt.Println("invoked TypeStr with", string(text))
	})

	s := hook.Start()
	<-hook.Process(s)
}

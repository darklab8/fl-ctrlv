package main

import (
	"fmt"
	"os"
	"strings"
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

		title_contains := "Freelancer,Wine"
		if value, ok := os.LookupEnv("TITLE_CONTAINS"); ok {
			title_contains = value
		}
		title_contains_split := strings.Split(title_contains, ",")
		active_title := robotgo.GetTitle()
		game_is_active := false

		for _, title := range title_contains_split {
			if strings.Contains(strings.ToLower(active_title), strings.ToLower(title)) {
				game_is_active = true
			}
		}
		if !game_is_active {
			fmt.Println("Game is not active. Targteting titles=", title_contains)
			return
		}

		fmt.Println("invoked TypeStr with", string(text))
		robotgo.TypeStr(string(text))
	})

	s := hook.Start()
	<-hook.Process(s)
}

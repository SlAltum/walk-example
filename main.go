package main

import (
	. "github.com/lxn/walk/declarative"
	"github.com/slaltum/walk-example/res"
	"github.com/slaltum/walk-example/src"
)

func main() {
	MainWindow{
		Title:   "WALK-EXAMPLE",
		Icon:    res.IconFavicon,
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "ACTION",
				OnClicked: func() {
					src.NewActionMainWindow().Run()
				},
			},
		},
	}.Run()
}

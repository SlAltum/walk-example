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
			PushButton{
				Text: "CLIPBOX",
				OnClicked: func() {
					src.NewClipBoardMainWindow().Run()
				},
			},
			PushButton{
				Text: "DATABINDING",
				OnClicked: func() {
					src.NewDataBindingMainWindow().Run()
				},
			},
			PushButton{
				Text: "DRAWING",
				OnClicked: func() {
					src.NewDrawingMainWindow().Run()
				},
			},
			PushButton{
				Text: "DROPFILE",
				OnClicked: func() {
					src.NewDropfileMainWindow().Run()
				},
			},
			PushButton{
				Text: "EXTERNAL_WIDGETS",
				OnClicked: func() {
					src.NewExternalWidgetsMainWindow().Run()
				},
			},
			PushButton{
				Text: "FILE_BROWSER",
				OnClicked: func() {
					src.NewFileBrowserMainWindow().Run()
				},
			},
			VSpacer{},
		},
	}.Run()
}

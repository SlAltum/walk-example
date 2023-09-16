package src

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewClipBoardMainWindow() *MainWindow {
	var te *walk.TextEdit

	mw := MainWindow{
		Icon:    res.IconFavicon,
		Title:   "Walk Clipboard Example",
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "Copy",
				OnClicked: func() {
					if err := walk.Clipboard().SetText(te.Text()); err != nil {
						log.Print("Copy: ", err)
					}
				},
			},
			PushButton{
				Text: "Paste",
				OnClicked: func() {
					if text, err := walk.Clipboard().Text(); err != nil {
						log.Print("Paste: ", err)
					} else {
						te.SetText(text)
					}
				},
			},
			TextEdit{
				AssignTo: &te,
			},
		},
	}
	return &mw
}

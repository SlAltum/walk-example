package src

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewDropfileMainWindow() *MainWindow {
	var textEdit *walk.TextEdit
	mw := MainWindow{
		Icon:    res.IconFavicon,
		Title:   "Walk DropFiles Example",
		MinSize: Size{320, 240},
		Layout:  VBox{},
		OnDropFiles: func(files []string) {
			textEdit.SetText(strings.Join(files, "\r\n"))
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &textEdit,
				ReadOnly: true,
				Text:     "Drop files here, from windows explorer...",
			},
		},
	}
	return &mw
}

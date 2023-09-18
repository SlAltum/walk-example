package src

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewLinkLabelMainWindow() *MainWindow {
	mw := MainWindow{
		Icon:    res.IconFavicon,
		Title:   "Walk LinkLabel Example",
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			LinkLabel{
				MaxSize: Size{100, 0},
				Text:    `I can contain multiple links like <a id="this" href="https://golang.org">this</a> or <a id="that" href="https://github.com/lxn/walk">that one</a>.`,
				OnLinkActivated: func(link *walk.LinkLabelLink) {
					log.Printf("id: '%s', url: '%s'\n", link.Id(), link.URL())
				},
			},
		},
	}
	return &mw
}

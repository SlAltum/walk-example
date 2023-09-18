package src

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func NewStatusBarAppRun() {
	icon1, err := walk.NewIconFromFile("res/img/check.ico")
	if err != nil {
		log.Fatal(err)
	}
	icon2, err := walk.NewIconFromFile("res/img/stop.ico")
	if err != nil {
		log.Fatal(err)
	}

	var sbi *walk.StatusBarItem

	MainWindow{
		Title:   "Walk Statusbar Example",
		MinSize: Size{600, 200},
		Layout:  VBox{MarginsZero: true},
		StatusBarItems: []StatusBarItem{
			{
				AssignTo: &sbi,
				Icon:     icon1,
				Text:     "click",
				Width:    80,
				OnClicked: func() {
					if sbi.Text() == "click" {
						sbi.SetText("again")
						sbi.SetIcon(icon2)
					} else {
						sbi.SetText("click")
						sbi.SetIcon(icon1)
					}
				},
			},
			{
				Text:        "left",
				ToolTipText: "no tooltip for me",
			},
			{
				Text: "\tcenter",
			},
			{
				Text: "\t\tright",
			},
			{
				Icon:        icon1,
				ToolTipText: "An icon with a tooltip",
			},
		},
	}.Run()
}

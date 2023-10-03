package externalwidgets

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

const myWidgetWindowClass = "MyWidget Class"

func init() {
	walk.AppendToWalkInit(func() {
		walk.MustRegisterWindowClass(myWidgetWindowClass)
	})
}

func NewExternalWidgetsMainWindow() *MainWindow {
	var mw *walk.MainWindow

	externalWidgetsMw := MainWindow{
		Icon:     res.IconFavicon,
		AssignTo: &mw,
		Title:    "Walk External Widgets Example",
		Size:     Size{Width: 400, Height: 300},
		Layout:   HBox{},
	}
	if err := externalWidgetsMw.Create(); err != nil {
		log.Fatal(err)
	}

	for _, name := range []string{"a", "b", "c"} {
		if w, err := NewMyWidget(mw); err != nil {
			log.Fatal(err)
		} else {
			w.SetName(name)
		}
	}

	mpb, err := NewMyPushButton(mw)
	if err != nil {
		log.Fatal(err)
	}
	mpb.SetText("MyPushButton")

	mw.Run()
	return &externalWidgetsMw
}

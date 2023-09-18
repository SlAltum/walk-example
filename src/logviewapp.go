package src

import (
	"log"
	"time"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func NewLogViewMainWindow() *MainWindow {
	var mw *walk.MainWindow

	logViewMw := MainWindow{
		AssignTo: &mw,
		Title:    "Walk LogView Example",
		MinSize:  Size{320, 240},
		Size:     Size{400, 600},
		Layout:   VBox{MarginsZero: true},
	}
	if err := (logViewMw.Create()); err != nil {
		log.Fatal(err)
	}

	lv, err := NewLogView(mw)
	if err != nil {
		log.Fatal(err)
	}

	lv.PostAppendText("XXX")
	log.SetOutput(lv)

	go func() {
		for i := 0; i < 10000; i++ {
			time.Sleep(100 * time.Millisecond)
			log.Println("Text" + "\r\n")
		}
	}()

	return &logViewMw
}

package databinding

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewDataBindingMainWindow() *MainWindow {
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})

	var mw *walk.MainWindow
	var outTE *walk.TextEdit

	animal := new(Animal)

	databingdingMw := MainWindow{
		AssignTo: &mw,
		Icon:     res.IconFavicon,
		Title:    "Walk Data Binding Example",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "Edit Animal",
				OnClicked: func() {
					if cmd, err := RunAnimalDialog(mw, animal); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						outTE.SetText(fmt.Sprintf("%+v", animal))
					}
				},
			},
			Label{
				Text: "animal:",
			},
			TextEdit{
				AssignTo: &outTE,
				ReadOnly: true,
				Text:     fmt.Sprintf("%+v", animal),
			},
		},
	}
	return &databingdingMw
}

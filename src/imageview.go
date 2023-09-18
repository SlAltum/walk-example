package src

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewImageViewMainWindow() *MainWindow {
	walk.Resources.SetRootDirPath("res/img")

	type Mode struct {
		Name  string
		Value ImageViewMode
	}

	modes := []Mode{
		{"ImageViewModeIdeal", ImageViewModeIdeal},
		{"ImageViewModeCorner", ImageViewModeCorner},
		{"ImageViewModeCenter", ImageViewModeCenter},
		{"ImageViewModeShrink", ImageViewModeShrink},
		{"ImageViewModeZoom", ImageViewModeZoom},
		{"ImageViewModeStretch", ImageViewModeStretch},
	}

	var widgets []Widget

	for _, mode := range modes {
		widgets = append(widgets,
			Label{
				Text: mode.Name,
			},
			ImageView{
				Background: SolidColorBrush{Color: walk.RGB(255, 191, 0)},
				Image:      "open.png",
				Margin:     10,
				Mode:       mode.Value,
			},
		)
	}

	mw := MainWindow{
		Icon:     res.IconFavicon,
		Title:    "Walk ImageView Example",
		Size:     Size{400, 600},
		Layout:   Grid{Columns: 2},
		Children: widgets,
	}
	return &mw
}

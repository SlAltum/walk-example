package externalwidgets

import (
	"log"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyWidget struct {
	walk.WidgetBase
}

func NewMyWidget(parent walk.Container) (*MyWidget, error) {
	w := new(MyWidget)

	if err := walk.InitWidget(
		w,
		parent,
		myWidgetWindowClass,
		win.WS_VISIBLE,
		0); err != nil {

		return nil, err
	}

	bg, err := walk.NewSolidColorBrush(walk.RGB(0, 255, 0))
	if err != nil {
		return nil, err
	}
	w.SetBackground(bg)

	return w, nil
}

func (*MyWidget) CreateLayoutItem(ctx *walk.LayoutContext) walk.LayoutItem {
	return &myWidgetLayoutItem{idealSize: walk.SizeFrom96DPI(walk.Size{50, 50}, ctx.DPI())}
}

func (w *MyWidget) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", w.Name())
	}

	return w.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

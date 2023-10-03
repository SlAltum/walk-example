package externalwidgets

import (
	"log"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyPushButton struct {
	*walk.PushButton
}

func NewMyPushButton(parent walk.Container) (*MyPushButton, error) {
	pb, err := walk.NewPushButton(parent)
	if err != nil {
		return nil, err
	}

	mpb := &MyPushButton{pb}

	if err := walk.InitWrapperWindow(mpb); err != nil {
		return nil, err
	}

	return mpb, nil
}

func (mpb *MyPushButton) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", mpb.Text())
	}

	return mpb.PushButton.WndProc(hwnd, msg, wParam, lParam)
}

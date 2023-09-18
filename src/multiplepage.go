package src

import (
	"bytes"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func NewMultiplepageAppMainWindow() *MultiplepageAppMainWindow {
	walk.Resources.SetRootDirPath("res/img")

	mw := new(MultiplepageAppMainWindow)

	cfg := &MultiPageMainWindowConfig{
		Name:    "mainWindow",
		MinSize: Size{600, 400},
		MenuItems: []MenuItem{
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "About",
						OnTriggered: func() { mw.aboutAction_Triggered() },
					},
				},
			},
		},
		OnCurrentPageChanged: func() {
			mw.updateTitle(mw.CurrentPageTitle())
		},
		PageCfgs: []PageConfig{
			{"Foo", "document-new.png", newFooPage},
			{"Bar", "document-properties.png", newBarPage},
			{"Baz", "system-shutdown.png", newBazPage},
		},
	}

	mpmw, err := NewMultiPageMainWindow(cfg)
	if err != nil {
		panic(err)
	}

	mw.MultiPageMainWindow = mpmw

	mw.updateTitle(mw.CurrentPageTitle())

	return mw
}

type MultiplepageAppMainWindow struct {
	*MultiPageMainWindow
}

func (mw *MultiplepageAppMainWindow) updateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString("Walk Multiple Pages Example")

	mw.SetTitle(buf.String())
}

func (mw *MultiplepageAppMainWindow) aboutAction_Triggered() {
	walk.MsgBox(mw,
		"About Walk Multiple Pages Example",
		"An example that demonstrates a main window that supports multiple pages.",
		walk.MsgBoxOK|walk.MsgBoxIconInformation)
}

type FooPage struct {
	*walk.Composite
}

func newFooPage(parent walk.Container) (Page, error) {
	p := new(FooPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "fooPage",
		Layout:   HBox{},
		Children: []Widget{
			HSpacer{},
			Label{Text: "I'm the Foo page"},
			HSpacer{},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

type BarPage struct {
	*walk.Composite
}

func newBarPage(parent walk.Container) (Page, error) {
	p := new(BarPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "barPage",
		Layout:   HBox{},
		Children: []Widget{
			HSpacer{},
			Label{Text: "I'm the Bar page"},
			HSpacer{},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

type BazPage struct {
	*walk.Composite
}

func newBazPage(parent walk.Container) (Page, error) {
	p := new(BazPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "bazPage",
		Layout:   HBox{},
		Children: []Widget{
			HSpacer{},
			Label{Text: "I'm the Baz page"},
			HSpacer{},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

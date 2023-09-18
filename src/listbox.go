package src

import (
	"fmt"
	"os"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewListBoxMainWindow() *MainWindow {
	mw := &ListBoxMainWindow{model: NewEnvModel()}

	listBoxMw := MainWindow{
		AssignTo: &mw.MainWindow,
		Icon:     res.IconFavicon,
		Title:    "Walk ListBox Example",
		MinSize:  Size{240, 320},
		Size:     Size{300, 400},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{
						AssignTo:              &mw.lb,
						Model:                 mw.model,
						OnCurrentIndexChanged: mw.lb_CurrentIndexChanged,
						OnItemActivated:       mw.lb_ItemActivated,
					},
					TextEdit{
						AssignTo: &mw.te,
						ReadOnly: true,
					},
				},
			},
		},
	}
	return &listBoxMw
}

type ListBoxMainWindow struct {
	*walk.MainWindow
	model *EnvModel
	lb    *walk.ListBox
	te    *walk.TextEdit
}

func (mw *ListBoxMainWindow) lb_CurrentIndexChanged() {
	i := mw.lb.CurrentIndex()
	item := &mw.model.items[i]

	mw.te.SetText(item.value)

	fmt.Println("CurrentIndex: ", i)
	fmt.Println("CurrentEnvVarName: ", item.name)
}

func (mw *ListBoxMainWindow) lb_ItemActivated() {
	value := mw.model.items[mw.lb.CurrentIndex()].value

	walk.MsgBox(mw, "Value", value, walk.MsgBoxIconInformation)
}

type EnvItem struct {
	name  string
	value string
}

type EnvModel struct {
	walk.ListModelBase
	items []EnvItem
}

func NewEnvModel() *EnvModel {
	env := os.Environ()

	m := &EnvModel{items: make([]EnvItem, len(env))}

	for i, e := range env {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		name := e[0:j]
		value := strings.Replace(e[j+1:], ";", "\r\n", -1)

		m.items[i] = EnvItem{name, value}
	}

	return m
}

func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index].name
}

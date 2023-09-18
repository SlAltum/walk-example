package main

import (
	. "github.com/lxn/walk/declarative"
	"github.com/slaltum/walk-example/res"
	"github.com/slaltum/walk-example/src"
)

func main() {
	MainWindow{
		Title:   "WALK-EXAMPLE",
		Icon:    res.IconFavicon,
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "ACTION",
				OnClicked: func() {
					src.NewActionMainWindow().Run()
				},
			},
			PushButton{
				Text: "CLIPBOX",
				OnClicked: func() {
					src.NewClipBoardMainWindow().Run()
				},
			},
			PushButton{
				Text: "DATABINDING",
				OnClicked: func() {
					src.NewDataBindingMainWindow().Run()
				},
			},
			PushButton{
				Text: "DRAWING",
				OnClicked: func() {
					src.NewDrawingMainWindow().Run()
				},
			},
			PushButton{
				Text: "DROPFILE",
				OnClicked: func() {
					src.NewDropfileMainWindow().Run()
				},
			},
			PushButton{
				Text: "EXTERNAL_WIDGETS",
				OnClicked: func() {
					src.NewExternalWidgetsMainWindow().Run()
				},
			},
			PushButton{
				Text: "FILE_BROWSER",
				OnClicked: func() {
					src.NewFileBrowserMainWindow().Run()
				},
			},
			PushButton{
				Text: "GRADIENT_COMPOSITE",
				OnClicked: func() {
					src.NewGradientCompositeMainWindow().Run()
				},
			},
			PushButton{
				Text: "IMAGE_ICON",
				OnClicked: func() {
					src.NewImageIconMainWindow().Run()
				},
			},
			PushButton{
				Text: "IMAGE_VIEW",
				OnClicked: func() {
					src.NewImageViewMainWindow().Run()
				},
			},
			PushButton{
				Text: "IMAGE_VIEWER",
				OnClicked: func() {
					src.NewImageViewerMainWindow().Run()
				},
			},
			PushButton{
				Text: "LINK_LABEL",
				OnClicked: func() {
					src.NewLinkLabelMainWindow().Run()
				},
			},
			PushButton{
				Text: "LIST_BOX",
				OnClicked: func() {
					src.NewListBoxMainWindow().Run()
				},
			},
			PushButton{
				Text: "LIST_BOX_OWNER_DRAWING",
				OnClicked: func() {
					src.NewListBoxOwnerDrawingMainWindow().Run()
					src.ListBoxOwnerDrawingCancel <- true
				},
			},
			PushButton{
				Text: "LOG_VIEW",
				OnClicked: func() {
					src.NewLogViewMainWindow().Run()
				},
			},
			PushButton{
				Text: "MULTIPLE_PAGE",
				OnClicked: func() {
					src.NewMultiplepageAppMainWindow().Run()
				},
			},
			PushButton{
				Text: "NOTIFY_ICON",
				OnClicked: func() {
					src.NewNotifyIconMainWindow().Run()
				},
			},
			VSpacer{},
		},
	}.Run()
}

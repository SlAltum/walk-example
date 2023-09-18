package src

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func NewSettingsAppRun() {
	app := walk.App()

	// These specify the app data sub directory for the settings file.
	app.SetOrganizationName("The Walk Authors")
	app.SetProductName("Walk Settings Example")

	// Settings file name.
	settings := walk.NewIniFileSettings("settings.ini")

	// All settings marked as expiring will expire after this duration w/o use.
	// This applies to all widgets settings.
	settings.SetExpireDuration(time.Hour * 24 * 30 * 3)

	if err := settings.Load(); err != nil {
		log.Fatal(err)
	}

	app.SetSettings(settings)

	if err := RunMainWindow(); err != nil {
		log.Fatal(err)
	}

	if err := settings.Save(); err != nil {
		log.Fatal(err)
	}
}

func RunMainWindow() error {
	if _, err := (MainWindow{
		Name:    "mainWindow", // Name is needed for settings persistence
		Title:   "Walk Settings Example",
		MinSize: Size{800, 600},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			TableView{
				Name:             "tableView", // Name is needed for settings persistence
				AlternatingRowBG: true,
				ColumnsOrderable: true,
				Columns: []TableViewColumn{
					// Name is needed for settings persistence
					{Name: "#", DataMember: "Index"}, // Use DataMember, if names differ
					{Name: "Bar"},
					{Name: "Baz", Format: "%.2f", Alignment: AlignFar},
					{Name: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
				},
				Model: NewSettingsFooModel(),
			}},
	}.Run()); err != nil {
		return err
	}

	return nil
}

func NewSettingsFooModel() *SettingsFooModel {
	now := time.Now()

	rand.New(rand.NewSource(now.UnixNano()))

	m := &SettingsFooModel{items: make([]*SettingsFoo, 1000)}

	for i := range m.items {
		m.items[i] = &SettingsFoo{
			Index: i,
			Bar:   strings.Repeat("*", rand.Intn(5)+1),
			Baz:   rand.Float64() * 1000,
			Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
		}
	}

	return m
}

type SettingsFooModel struct {
	walk.SortedReflectTableModelBase
	items []*SettingsFoo
}

func (m *SettingsFooModel) Items() interface{} {
	return m.items
}

type SettingsFoo struct {
	Index int
	Bar   string
	Baz   float64
	Quux  time.Time
}

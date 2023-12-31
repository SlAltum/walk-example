package databinding

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func RunAnimalDialog(owner walk.Form, animal *Animal) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         Bind("'Animal Details' + (animal.Name == '' ? '' : ' - ' + animal.Name)"),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "animal",
			DataSource:     animal,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Name:",
					},
					LineEdit{
						Text: Bind("Name"),
					},

					Label{
						Text: "Arrival Date:",
					},
					DateEdit{
						Date: Bind("ArrivalDate"),
					},

					Label{
						Text: "Species:",
					},
					ComboBox{
						Value:         Bind("SpeciesId", SelRequired{}),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         KnownSpecies(),
					},

					Label{
						Text: "Speed:",
					},
					Slider{
						Value: Bind("Speed"),
					},

					RadioButtonGroupBox{
						ColumnSpan: 2,
						Title:      "Sex",
						Layout:     HBox{},
						DataMember: "Sex",
						Buttons: []RadioButton{
							{Text: "Male", Value: SexMale},
							{Text: "Female", Value: SexFemale},
							{Text: "Hermaphrodite", Value: SexHermaphrodite},
						},
					},

					Label{
						Text: "Weight:",
					},
					NumberEdit{
						Value:    Bind("Weight", Range{0.01, 9999.99}),
						Suffix:   " kg",
						Decimals: 2,
					},

					Label{
						Text: "Preferred Food:",
					},
					ComboBox{
						Editable: true,
						Value:    Bind("PreferredFood"),
						Model:    []string{"Fruit", "Grass", "Fish", "Meat"},
					},

					Label{
						Text: "Domesticated:",
					},
					CheckBox{
						Checked: Bind("Domesticated"),
					},

					VSpacer{
						ColumnSpan: 2,
						Size:       8,
					},

					Label{
						ColumnSpan: 2,
						Text:       "Remarks:",
					},
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},

					Label{
						ColumnSpan: 2,
						Text:       "Patience:",
					},
					LineEdit{
						ColumnSpan: 2,
						Text:       Bind("PatienceField"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

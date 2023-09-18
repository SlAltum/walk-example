package src

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
	"github.com/slaltum/walk-example/res"
)

type TableViewFoo struct {
	Index   int
	Bar     string
	Baz     float64
	Quux    time.Time
	checked bool
}

type TableViewFooModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*TableViewFoo
}

func NewTableViewFooModel() *TableViewFooModel {
	m := new(TableViewFooModel)
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *TableViewFooModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *TableViewFooModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.Bar

	case 2:
		return item.Baz

	case 3:
		return item.Quux
	}

	panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *TableViewFooModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *TableViewFooModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Called by the TableView to sort the model.
func (m *TableViewFooModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.Index < b.Index)

		case 1:
			return c(a.Bar < b.Bar)

		case 2:
			return c(a.Baz < b.Baz)

		case 3:
			return c(a.Quux.Before(b.Quux))
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *TableViewFooModel) ResetRows() {
	// Create some random data.
	m.items = make([]*TableViewFoo, rand.Intn(50000))

	now := time.Now()

	for i := range m.items {
		m.items[i] = &TableViewFoo{
			Index: i,
			Bar:   strings.Repeat("*", rand.Intn(5)+1),
			Baz:   rand.Float64() * 1000,
			Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

func RunNewTableViewApp() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	boldFont, _ := walk.NewFont("Segoe UI", 9, walk.FontBold)
	goodIcon, _ := walk.Resources.Icon("../img/check.ico")
	badIcon, _ := walk.Resources.Icon("../img/stop.ico")

	barBitmap, err := walk.NewBitmapForDPI(walk.Size{100, 1}, res.DPI)
	if err != nil {
		panic(err)
	}
	defer barBitmap.Dispose()

	canvas, err := walk.NewCanvasFromImage(barBitmap)
	if err != nil {
		panic(err)
	}
	defer barBitmap.Dispose()

	canvas.GradientFillRectanglePixels(walk.RGB(255, 0, 0), walk.RGB(0, 255, 0), walk.Horizontal, walk.Rectangle{0, 0, 100, 1})

	canvas.Dispose()

	model := NewTableViewFooModel()

	var tv *walk.TableView

	MainWindow{
		Title:  "Walk TableView Example",
		Size:   Size{800, 600},
		Layout: VBox{MarginsZero: true},
		Children: []Widget{
			PushButton{
				Text:      "Reset Rows",
				OnClicked: model.ResetRows,
			},
			PushButton{
				Text: "Select first 5 even Rows",
				OnClicked: func() {
					tv.SetSelectedIndexes([]int{0, 2, 4, 6, 8})
				},
			},
			TableView{
				AssignTo:         &tv,
				AlternatingRowBG: true,
				CheckBoxes:       true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "Bar"},
					{Title: "Baz", Alignment: AlignFar},
					{Title: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
				},
				StyleCell: func(style *walk.CellStyle) {
					item := model.items[style.Row()]

					if item.checked {
						if style.Row()%2 == 0 {
							style.BackgroundColor = walk.RGB(159, 215, 255)
						} else {
							style.BackgroundColor = walk.RGB(143, 199, 239)
						}
					}

					switch style.Col() {
					case 1:
						if canvas := style.Canvas(); canvas != nil {
							bounds := style.Bounds()
							bounds.X += 2
							bounds.Y += 2
							bounds.Width = int((float64(bounds.Width) - 4) / 5 * float64(len(item.Bar)))
							bounds.Height -= 4
							canvas.DrawBitmapPartWithOpacityPixels(barBitmap, bounds, walk.Rectangle{0, 0, 100 / 5 * len(item.Bar), 1}, 127)

							bounds.X += 4
							bounds.Y += 2
							canvas.DrawText(item.Bar, tv.Font(), 0, bounds, walk.TextLeft)
						}

					case 2:
						if item.Baz >= 900.0 {
							style.TextColor = walk.RGB(0, 191, 0)
							style.Image = goodIcon
						} else if item.Baz < 100.0 {
							style.TextColor = walk.RGB(255, 0, 0)
							style.Image = badIcon
						}

					case 3:
						if item.Quux.After(time.Now().Add(-365 * 24 * time.Hour)) {
							style.Font = boldFont
						}
					}
				},
				Model: model,
				OnSelectedIndexesChanged: func() {
					fmt.Printf("SelectedIndexes: %v\n", tv.SelectedIndexes())
				},
			},
		},
	}.Run()
}

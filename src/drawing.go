package src

import (
	"math"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/slaltum/walk-example/res"
)

func NewDrawingMainWindow() *MainWindow {
	mw := new(DrawingMainWindow)

	drawingMw := MainWindow{
		AssignTo: &mw.MainWindow,
		Icon:     res.IconFavicon,
		Title:    "Walk Drawing Example",
		MinSize:  Size{320, 240},
		Size:     Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			CustomWidget{
				AssignTo:            &mw.paintWidget,
				ClearsBackground:    true,
				InvalidatesOnResize: true,
				Paint:               mw.drawStuff,
			},
		},
	}
	return &drawingMw
}

type DrawingMainWindow struct {
	*walk.MainWindow
	paintWidget *walk.CustomWidget
}

func (mw *DrawingMainWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bmp, err := createBitmap()
	if err != nil {
		return err
	}
	defer bmp.Dispose()

	bounds := mw.paintWidget.ClientBounds()

	rectPen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(255, 0, 0))
	if err != nil {
		return err
	}
	defer rectPen.Dispose()

	if err := canvas.DrawRectanglePixels(rectPen, bounds); err != nil {
		return err
	}

	ellipseBrush, err := walk.NewHatchBrush(walk.RGB(0, 255, 0), walk.HatchCross)
	if err != nil {
		return err
	}
	defer ellipseBrush.Dispose()

	if err := canvas.FillEllipsePixels(ellipseBrush, bounds); err != nil {
		return err
	}

	linesBrush, err := walk.NewSolidColorBrush(walk.RGB(0, 0, 255))
	if err != nil {
		return err
	}
	defer linesBrush.Dispose()

	linesPen, err := walk.NewGeometricPen(walk.PenDash, 8, linesBrush)
	if err != nil {
		return err
	}
	defer linesPen.Dispose()

	if err := canvas.DrawLinePixels(linesPen, walk.Point{bounds.X, bounds.Y}, walk.Point{bounds.Width, bounds.Height}); err != nil {
		return err
	}
	if err := canvas.DrawLinePixels(linesPen, walk.Point{bounds.X, bounds.Height}, walk.Point{bounds.Width, bounds.Y}); err != nil {
		return err
	}

	points := make([]walk.Point, 10)
	dx := bounds.Width / (len(points) - 1)
	for i := range points {
		points[i].X = i * dx
		points[i].Y = int(float64(bounds.Height) / math.Pow(float64(bounds.Width/2), 2) * math.Pow(float64(i*dx-bounds.Width/2), 2))
	}
	for i, p := range points {
		if err := canvas.DrawLinePixels(linesPen, p, points[(i+1)%len(points)]); err != nil {
			return err
		}
	}

	bmpSize := bmp.Size()
	if err := canvas.DrawImagePixels(bmp, walk.Point{(bounds.Width - bmpSize.Width) / 2, (bounds.Height - bmpSize.Height) / 2}); err != nil {
		return err
	}

	return nil
}

func createBitmap() (*walk.Bitmap, error) {
	bounds := walk.Rectangle{Width: 200, Height: 200}

	bmp, err := walk.NewBitmapForDPI(bounds.Size(), res.DPI)
	if err != nil {
		return nil, err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			bmp.Dispose()
		}
	}()

	canvas, err := walk.NewCanvasFromImage(bmp)
	if err != nil {
		return nil, err
	}
	defer canvas.Dispose()

	brushBmp, err := walk.NewBitmapFromFileForDPI("res/img/plus.png", res.DPI)
	if err != nil {
		return nil, err
	}
	defer brushBmp.Dispose()

	brush, err := walk.NewBitmapBrush(brushBmp)
	if err != nil {
		return nil, err
	}
	defer brush.Dispose()

	if err := canvas.FillRectanglePixels(brush, bounds); err != nil {
		return nil, err
	}

	font, err := walk.NewFont("Times New Roman", 40, walk.FontBold|walk.FontItalic)
	if err != nil {
		return nil, err
	}
	defer font.Dispose()

	if err := canvas.DrawTextPixels("Walk Drawing Example", font, walk.RGB(0, 0, 0), bounds, walk.TextWordbreak); err != nil {
		return nil, err
	}

	succeeded = true

	return bmp, nil
}

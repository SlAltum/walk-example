package externalwidgets

import (
	"github.com/lxn/walk"
)

type myWidgetLayoutItem struct {
	walk.LayoutItemBase
	idealSize walk.Size // in native pixels
}

func (li *myWidgetLayoutItem) LayoutFlags() walk.LayoutFlags {
	return 0
}

func (li *myWidgetLayoutItem) IdealSize() walk.Size {
	return li.idealSize
}

package UI

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func HelpWidget() ui.Drawable {
	l := widgets.NewList()

	l.Rows = GetCommandDesc()
	l.SetRect(0, 0, width, height)
	return l
}

package UI

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"oraTop/db"
)

func OraSessTable() ui.Drawable {
	data := db.OraDB.GetSession()
	tab := widgets.NewTable()
	tab.Rows = data
	tab.SetRect(0, 0, width, height)
	return tab
}

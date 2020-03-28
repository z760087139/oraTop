package UI

import (
	ui "github.com/gizak/termui/v3"
)

var (
	height int
	width  int
)

func GetTermSize() {
	width, height = ui.TerminalDimensions()
}

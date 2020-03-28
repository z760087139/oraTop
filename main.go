package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"oraTop/db"
	UI "oraTop/widget"
)

func makeUI() {
	// 启动UI
	if err := ui.Init(); err != nil {
		log.Fatal("create termUI failed")
	}
	defer ui.Close()
	// 获取widget 列表
	s := `ss`
	w := widgets.NewParagraph()
	w.Text = s
	w.SetRect(0, 0, 40, 40)
	ui.Render(w)

	e := ui.PollEvents()
	for {
		t := <-e
		switch t.ID {
		case "q", "<C-c>":
			return
		}
	}
}

func main() {
	defer db.OraDB.Close()
	if err := ui.Init(); err != nil {
		log.Fatal("create termUI failed")
	}
	defer ui.Close()
	// 初始化ui之后才能调用
	UI.GetTermSize()

	ui.Render(UI.HelpWidget())

	var currentWidget func() ui.Drawable
	currentWidget = UI.HelpWidget

	// 事件采集及命令判断 定时刷新
	e := ui.PollEvents()
	ticker := time.NewTicker(time.Second * 2).C
	for {
		select {
		case t := <-e:
			if t.Type == ui.ResizeEvent {
				UI.GetTermSize()
				continue
			} else if t.Type == ui.MouseEvent {
				continue
			}
			// keyboard input
			command, ok := UI.CommandMap[t.ID]
			if !ok {
				continue
			}
			if t.ID == "q" || t.ID == "<C-c>" {
				return
			}
			currentWidget = command.Widget
			ui.Render(currentWidget())
		case <-ticker:
			ui.Render(currentWidget())
		}
	}
}

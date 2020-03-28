package UI

import (
	"fmt"
	"sort"

	ui "github.com/gizak/termui/v3"
)

var (
	CommandMap map[string]TopCommand
)

type TopCommand struct {
	Command string
	Desc    string
	Widget  func() ui.Drawable
}

func init() {
	CommandMap = map[string]TopCommand{
		"s": TopCommand{
			Desc:   `Current Session, data from v$session`,
			Widget: OraSessTable,
		},
		"h": TopCommand{
			Desc:   `Help Info`,
			Widget: HelpWidget,
		},
		"q": TopCommand{
			Desc: `Quit`,
		},
	}
}

func GetCommandDesc() (commandDescList []string) {
	l := make([]string, 0, len(CommandMap))
	for k := range CommandMap {
		l = append(l, k)
	}
	sort.Strings(l)
	commandDescFmt := `[%s] : %s`
	commandDescList = make([]string, 0, len(CommandMap))
	for _, v := range l {
		commandDesc := fmt.Sprintf(commandDescFmt, v, CommandMap[v].Desc)
		commandDescList = append(commandDescList, commandDesc)
	}
	return
}

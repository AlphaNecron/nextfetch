package nextfetch

import (
	"fmt"
	"github.com/gookit/color"
	"nextfetch/src/nextfetch/utils"
)

type Info struct {
	Name  string
	Value string
}

func (i Info) print(indentSize int) {
	if !config.NoAscii {
		utils.MoveCursorForward(indentSize)
	}
	color.Print(fmt.Sprintf("<fg=%s>%s</>: <fg=white>%s</>\n", config.PrimaryColor, i.Name, i.Value))
}

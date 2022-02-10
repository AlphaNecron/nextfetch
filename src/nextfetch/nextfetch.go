package nextfetch

import (
	"fmt"
	"github.com/gookit/color"
	"nextfetch/src/nextfetch/ascii"
	"nextfetch/src/nextfetch/constants"
	"nextfetch/src/nextfetch/fetcher"
	"nextfetch/src/nextfetch/utils"
	"strings"
)

func Nextfetch() {
	if config.ClearScreen {
		fmt.Print(constants.SCREEN_ERASER)
	}
	artWidth := 0
	os := fetcher.GetOs()
	if !config.NoAscii {
		art := ascii.ARTS["Generic"]
		art.Color = config.PrimaryColor
		if a, ok := ascii.ARTS[os]; ok {
			art = a
		}
		color.Println(fmt.Sprintf("<fg=%s>%s</>", art.Color, art.Art))
		utils.MoveCursorUp(utils.CountLine(art.Art))
		artWidth = utils.GetLongestWidth(strings.Split(art.Art, "\n")) + 2
		utils.MoveCursorForward(artWidth)
	}
	info := [6]Info{
		{"os", os},
		{"arch", fetcher.GetArch()},
		{"kernel", fetcher.GetKernel()},
		{"uptime", fetcher.GetUptime()},
		{"cpu", utils.ShortenCpu(fetcher.GetCpuBrand())},
		{"mem", fetcher.GetMemInfo()},
	}
	color.Println(fmt.Sprintf("<fg=%[1]s;op=bold>%[2]s</><white;op=bold>@</><fg=%[1]s;op=bold>%[3]s</>", config.PrimaryColor, fetcher.GetUsername(), fetcher.GetHostname()))
	for _, inf := range info {
		inf.print(artWidth)
	}
	utils.MoveCursorForward(artWidth)
	color.Print(utils.MakeColorBlock(config.ColorBlock))
}

package utils

import (
	"fmt"
	_ "image/png"
	"io/ioutil"
	"log"
	"nextfetch/src/nextfetch/constants"
	"regexp"
	"strconv"
	"strings"
)

func MakeColorBlock(bchar string) string {
	var block string
	for i := 0; i < 8; i++ {
		block += fmt.Sprintf("<fg=%s>%s</>", strconv.Itoa(i), bchar)
	}
	block += strings.Repeat(constants.Linebreak, 2)
	return block
}

func GetLongestWidth(arr []string) int {
	var max = len(arr[0])
	for _, str := range arr {
		if len(str) > max {
			max = len(str)
		}
	}
	return max
}

func CountLine(s string) int {
	n := strings.Count(s, "\n")
	if len(s) > 0 && !strings.HasSuffix(s, "\n") {
		n++
	}
	return n
}

// goterm
func MoveCursorForward(bias int) {
	fmt.Printf("\033[%dC", bias)
}

func MoveCursorUp(bias int) {
	fmt.Printf("\033[%dA", bias)
}

func ShortenCpu(cpu string) string {
	r := regexp.MustCompile(`(?i)(\((R|TM))\)|@\s[0-9]\.[0-9]{2}[MG]Hz$|Core|Core? Duo|CPU|Intel|AMD|Qualcomm|Processor|(Dual|Quad|Six|Eight|[0-9]+)-Core`)
	return strings.Replace(strings.TrimSpace(r.ReplaceAllString(cpu, "")), "  ", " ", 1)
}

func ReadLine(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content), constants.Linebreak)
}

func FormatUnit(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

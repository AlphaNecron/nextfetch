package nextfetch

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

var config Config

type Config struct {
	NoAscii      bool   `json:"noAscii"`
	PrimaryColor string `json:"primaryColor"`
	ColorBlock   string `json:"colorBlock"`
	ClearScreen  bool   `json:"clearScreen"`
}

func setDefConf() {
	config = Config{
		NoAscii:      false,
		PrimaryColor: "blue",
		ColorBlock:   "\xe2\x96\x81\xe2\x96\x81",
		ClearScreen:  false,
	}
}

func TryRead() {
	hd, err := os.UserHomeDir()
	hd = path.Join(hd, ".nextfetch")
	if err != nil {
		setDefConf()
		return
	}
	jF, err := os.Open(hd)
	if err != nil {
		setDefConf()
		return
	}
	defer func(jF *os.File) {
		err := jF.Close()
		if err != nil {
			setDefConf()
			return
		}
	}(jF)
	b, err := ioutil.ReadAll(jF)
	if err != nil {
		setDefConf()
		return
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		setDefConf()
		return
	}
}

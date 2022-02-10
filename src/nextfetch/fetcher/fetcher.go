package fetcher

import (
	"log"
	"os"
	"os/user"
	"strings"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return hostname
}

func GetUsername() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err.Error())
	}
	s := strings.Split(usr.Username, "\\")
	return s[len(s)-1]
}

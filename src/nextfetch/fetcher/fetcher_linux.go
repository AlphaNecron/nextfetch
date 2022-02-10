package fetcher

import (
	"fmt"
	"github.com/dekobon/distro-detect/linux"
	"golang.org/x/sys/unix"
	"log"
	"nextfetch/src/nextfetch/constants"
	"nextfetch/src/nextfetch/utils"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetOs() string {
	return linux.DiscoverDistro().Name
}

func GetKernel() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func getSysInfo() *unix.Sysinfo_t {
	sysInfo := &unix.Sysinfo_t{}
	if err := unix.Sysinfo(sysInfo); err != nil {
		log.Fatal(err)
	}
	return sysInfo
}

func GetUptime() string {
	return (time.Duration(getSysInfo().Uptime) * time.Second).String()
}

func GetCpuBrand() string {
	lines := utils.ReadLine(constants.PROC_CPUINFO)
	sort.Strings(lines)
	n := sort.SearchStrings(lines, "model name")
	model := lines[n]
	if !strings.HasPrefix(model, "model name") {
		return "unknown"
	}
	model = strings.TrimSpace(strings.Split(model, ":")[1])
	return model
}

func GetMemInfo() string { // gopsutil
	lines := utils.ReadLine(constants.PROC_MEMINFO)
	var memCached, memBuffer, memFree, sReclaimable, memTotal uint64
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)
		switch key {
		case "MemTotal":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			memTotal = t * 1024
		case "MemFree":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			memFree = t * 1024
		case "Buffers":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			memBuffer = t * 1024
		case "Cached":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			memCached = t * 1024
		case "SReclaimable":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			sReclaimable = t * 1024
		}
	}
	memCached += sReclaimable
	return fmt.Sprintf("%s / %s", utils.FormatUnit(memTotal-memCached-memBuffer-memFree), utils.FormatUnit(memTotal))
}

func GetArch() string {
	out, err := exec.Command("uname", "-m").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

package fetcher

import (
	"fmt"
	"github.com/yusufpapurcu/wmi"
	"golang.org/x/sys/windows/registry"
	"log"
	"nextfetch/src/nextfetch/utils"
	"syscall"
	"time"
	"unsafe"
)

var k32 = syscall.NewLazyDLL("kernel32.dll")

func GetOs() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		log.Fatal(err)
	}
	return pn
}

func GetUptime() string { // took from gopsutil
	procGetTickCount64 := k32.NewProc("GetTickCount64")
	procGetTickCount := procGetTickCount64
	err := procGetTickCount64.Find()
	if err != nil {
		procGetTickCount = k32.NewProc("GetTickCount")
	}
	r1, _, e := syscall.Syscall(procGetTickCount.Addr(), 0, 0, 0, 0)
	if e != 0 {
		return "unknown"
	}
	return (time.Duration(r1) * time.Millisecond).Truncate(time.Second).String()
}

func GetCpuBrand() string {
	type win32_Processor struct {
		Name string
	}
	var cpu []win32_Processor
	q := wmi.CreateQuery(&cpu, "")
	if err := wmi.Query(q, &cpu); err != nil {
		log.Fatalln(err)
	}
	return cpu[0].Name
}

func GetArch() string { // took from gopsutil
	type systemInfo struct {
		wProcessorArchitecture uint16
		wProcessorLevel        uint16
	}
	var sysInfo systemInfo
	k32.NewProc("GetNativeSystemInfo").Call(uintptr(unsafe.Pointer(&sysInfo)))
	const (
		PROCESSOR_ARCHITECTURE_INTEL = 0
		PROCESSOR_ARCHITECTURE_ARM   = 5
		PROCESSOR_ARCHITECTURE_ARM64 = 12
		PROCESSOR_ARCHITECTURE_IA64  = 6
		PROCESSOR_ARCHITECTURE_AMD64 = 9
	)
	switch sysInfo.wProcessorArchitecture {
	case PROCESSOR_ARCHITECTURE_INTEL:
		if sysInfo.wProcessorLevel < 3 {
			return "i386"
		}
		if sysInfo.wProcessorLevel > 6 {
			return "i686"
		}
		return fmt.Sprintf("i%d86", sysInfo.wProcessorLevel)
	case PROCESSOR_ARCHITECTURE_ARM:
		return "arm"
	case PROCESSOR_ARCHITECTURE_ARM64:
		return "aarch64"
	case PROCESSOR_ARCHITECTURE_IA64:
		return "ia64"
	case PROCESSOR_ARCHITECTURE_AMD64:
		return "x86_64"
	}
	return "unknown"
}

func GetMemInfo() string {
	type memoryStatusEx struct {
		cbSize                  uint32
		dwMemoryLoad            uint32
		ullTotalPhys            uint64
		ullAvailPhys            uint64
		ullTotalPageFile        uint64
		ullAvailPageFile        uint64
		ullTotalVirtual         uint64
		ullAvailVirtual         uint64
		ullAvailExtendedVirtual uint64
	}
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := k32.NewProc("GlobalMemoryStatusEx").Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return "unknown"
	}
	return fmt.Sprintf("%s / %s", utils.FormatUnit(memInfo.ullTotalPhys-memInfo.ullAvailPhys), utils.FormatUnit(memInfo.ullTotalPhys))
}

func GetKernel() string {
	return "unknown"
}

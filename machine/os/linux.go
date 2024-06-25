package os

import (
	"fmt"
	"strings"

	"github.com/xishengcai/machine-code/machine/types"
)

type LinuxMachine struct{}

func (i LinuxMachine) GetMachine() (types.Information, error) {
	platformUUID, err := i.GetPlatformUUID()
	if err != nil {
		return types.Information{}, err
	}
	boardSerialNumber, err := i.GetBoardSerialNumber()
	if err != nil {
		return types.Information{}, err
	}
	cpuSerialNumber, err := i.GetCpuSerialNumber()
	if err != nil {
		return types.Information{}, err
	}

	machineData := types.Information{
		PlatformUUID:      platformUUID,
		BoardSerialNumber: boardSerialNumber,
		CpuSerialNumber:   cpuSerialNumber,
	}
	return machineData, err
}

func (LinuxMachine) GetBoardSerialNumber() (serialNumber string, err error) {
	out, err := RunShellCommand("dmidecode -s system-serial-number")
	if err != nil {
		return "", fmt.Errorf("获取Linux 系统序列号失败, %s", err.Error())
	}
	fmt.Println("out:", out)
	serialNumber = strings.Replace(out, "\n", "", -1)
	return
}

func (LinuxMachine) GetPlatformUUID() (UUID string, err error) {
	out, err := RunShellCommand("dmidecode -s system-uuid")
	if err != nil {
		return "", fmt.Errorf("获取Linux 平台UUID失败, %s", err.Error())
	}
	fmt.Println("out:", out)
	UUID = strings.Replace(out, "\n", "", -1)
	return
}

func (LinuxMachine) GetCpuSerialNumber() (cpuId string, err error) {

	out, err := RunShellCommand("dmidecode -t processor |grep ID |head -1")
	if err != nil {
		return "", fmt.Errorf("获取Linux CPU ID 失败, %s", err.Error())
	}
	fmt.Println("out:", out)
	cpuId = strings.Replace(out, "\n", "", -1)
	return
}

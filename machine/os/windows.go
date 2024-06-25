package os

import (
	"os/exec"
	"strings"

	"github.com/xishengcai/machine-code/machine/types"
)

type WindowsMachine struct{}

func (i WindowsMachine) GetMachine() (types.Information, error) {
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

func (WindowsMachine) GetBoardSerialNumber() (serialNumber string, err error) {
	// wmic baseboard get serialnumber
	cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	b, e := cmd.CombinedOutput()
	if e == nil {
		serialNumber = string(b)
		serialNumber = serialNumber[12 : len(serialNumber)-2]
		serialNumber = strings.ReplaceAll(serialNumber, "\n", "")
		serialNumber = strings.ReplaceAll(serialNumber, " ", "")
		serialNumber = strings.ReplaceAll(serialNumber, "\r", "")
	} else {
		return "", nil
	}
	return serialNumber, nil
}

func (WindowsMachine) GetPlatformUUID() (uuid string, err error) {
	// wmic csproduct get uuid
	var cmd *exec.Cmd
	cmd = exec.Command("wmic", "csproduct", "get", "uuid")
	b, e := cmd.CombinedOutput()

	if e == nil {
		uuid = string(b)
		uuid = uuid[4 : len(uuid)-1]
		uuid = strings.ReplaceAll(uuid, "\n", "")
		uuid = strings.ReplaceAll(uuid, " ", "")
		uuid = strings.ReplaceAll(uuid, "\r", "")
	} else {
		return "", nil
	}
	return uuid, nil
}

func (WindowsMachine) GetCpuSerialNumber() (cpuId string, err error) {
	// wmic cpu get processorid
	var cpuid string
	cmd := exec.Command("wmic", "cpu", "get", "processorid")
	b, e := cmd.CombinedOutput()

	if e == nil {
		cpuid = string(b)
		cpuid = cpuid[12 : len(cpuid)-2]
		cpuid = strings.ReplaceAll(cpuid, "\n", "")
		cpuid = strings.ReplaceAll(cpuid, " ", "")
		cpuid = strings.ReplaceAll(cpuid, "\r", "")
	} else {
		return "", nil
	}
	return cpuid, nil
}

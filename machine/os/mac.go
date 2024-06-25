package os

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/xishengcai/machine-code/machine/types"
)

type MacMachine struct{}

var macMachineData types.Information

type macSysInfoStruct struct {
	SPHardwareDataType []struct {
		Name                 string `json:"_name"`
		ActivationLockStatus string `json:"activation_lock_status"`
		BootRomVersion       string `json:"boot_rom_version"`
		ChipType             string `json:"chip_type"`
		MachineModel         string `json:"machine_model"`
		MachineName          string `json:"machine_name"`
		ModelNumber          string `json:"model_number"`
		NumberProcessors     string `json:"number_processors"`
		OsLoaderVersion      string `json:"os_loader_version"`
		PhysicalMemory       string `json:"physical_memory"`
		PlatformUUID         string `json:"platform_UUID"`
		ProvisioningUDID     string `json:"provisioning_UDID"`
		SerialNumber         string `json:"serial_number"`
	} `json:"SPHardwareDataType"`
}

func (mac MacMachine) GetMachine() (types.Information, error) {
	platformUUID, err := mac.GetPlatformUUID()
	if err != nil {
		return types.Information{}, err
	}
	boardSerialNumber, err := mac.GetBoardSerialNumber()
	if err != nil {
		return types.Information{}, err
	}

	machineData := types.Information{
		PlatformUUID:      platformUUID,
		BoardSerialNumber: boardSerialNumber,
	}
	return machineData, nil
}

func (mac MacMachine) GetBoardSerialNumber() (data string, err error) {
	sysInfo, err := mac.GetMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.BoardSerialNumber, err
}

func (mac MacMachine) GetPlatformUUID() (UUID string, err error) {
	sysInfo, err := mac.GetMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.PlatformUUID, err
}

func (mac MacMachine) GetCpuSerialNumber() (cpuId string, err error) {
	sysInfo, err := mac.GetMacSysInfo()
	if err != nil {
		return "", err
	}
	return sysInfo.CpuSerialNumber, err
}

func (mac MacMachine) GetMacSysInfo() (data types.Information, err error) {
	out, err := RunShellCommand("system_profiler SPHardwareDataType -json")
	if err != nil {
		return data, fmt.Errorf("获取系统信息失败, %s", err.Error())
	}
	fmt.Println("out:", out)
	macMachineData, err = mac.macXmlToData(out)
	return macMachineData, err
}

func (MacMachine) macXmlToData(content string) (types.Information, error) {
	x := macSysInfoStruct{}
	err := json.Unmarshal([]byte(content), &x)
	if err != nil {
		return types.Information{}, err
	}
	serialData := types.Information{
		PlatformUUID:      x.SPHardwareDataType[0].PlatformUUID,
		BoardSerialNumber: x.SPHardwareDataType[0].SerialNumber,
		CpuSerialNumber:   "",
	}
	return serialData, nil

}

// RunShellCommand 执行shell命令并返回输出
func RunShellCommand(command string) (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

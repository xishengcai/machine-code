package machine

import "github.com/xishengcai/machine-code/machine/types"

// OsMachineInterface 机器码接口
type OsMachineInterface interface {
	GetMachine() (types.Information, error)
	GetBoardSerialNumber() (string, error)
	GetPlatformUUID() (string, error)
	GetCpuSerialNumber() (string, error)
}

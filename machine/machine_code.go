package machine

import (
	"errors"
	"net"
	"runtime"
	"strings"

	"github.com/xishengcai/machine-code/machine/os"
	"github.com/xishengcai/machine-code/machine/types"
)

func GetMachineData() (data types.Information) {
	var osMachine OsMachineInterface
	if runtime.GOOS == "darwin" {
		osMachine = os.MacMachine{}
	} else if runtime.GOOS == "linux" {
		osMachine = os.LinuxMachine{}
	} else if runtime.GOOS == "windows" {
		osMachine = os.WindowsMachine{}
	}
	machineData, err := osMachine.GetMachine()
	if err != nil {
		panic(err)
	}
	return machineData
}

func GetMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var mac string
	var bakMac1 string
	var bakMac2 string

	for i := 0; i < len(netInterfaces); i++ {
		flags := netInterfaces[i].Flags.String()

		if strings.Contains(flags, "up") &&
			strings.Contains(flags, "broadcast") &&
			strings.Contains(flags, "running") &&
			!strings.Contains(flags, "loopback") {

			//fmt.Println(fmt.Sprintf("i:%d name:%s %v", i, netInterfaces[i].Name, flags))
			if strings.Contains(netInterfaces[i].Name, "WLAN") {
				mac = netInterfaces[i].HardwareAddr.String()
				return mac, nil
			}
			if !strings.Contains(netInterfaces[i].Name, "VMware") {
				bakMac1 = netInterfaces[i].HardwareAddr.String()
			} else {
				bakMac2 = netInterfaces[i].HardwareAddr.String()
			}
		}
	}
	if mac == "" {
		if bakMac1 != "" {
			return bakMac1, nil
		}
		return bakMac2, nil
	}
	return mac, errors.New("unable to get the correct MAC address")
}

func GetLocalIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip, nil
}

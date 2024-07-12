package main

import (
	"encoding/json"
	"fmt"

	"github.com/xishengcai/machine-code/machine"
)

// https://www.icode9.com/content-3-710187.html  go 获取linux cpuId 的方法
func main() {
	machineData, err := machine.GetMachineData()
	if err != nil {
		panic(err)
	}
	result, err := json.Marshal(machineData)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(result))
}

package types

type Information struct {
	PlatformUUID      string `json:"platformUUID"`
	BoardSerialNumber string `json:"boardSerialNumber"` // 主板序列号  使用WMIC获取
	CpuSerialNumber   string `json:"cpuSerialNumber"`   // cpu序列号  使用WMIC获取
}

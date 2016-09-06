package main

// SystemInfo holds system information
type SystemInfo struct {
	SerialNumber string
	HardwareUUID string
}

// GetSystemInfo collects information about the system
func GetSystemInfo() (sysinfo SystemInfo) {
	sysinfo.SerialNumber = GetCommandOutput("system_profiler SPHardwareDataType | grep \"Serial Number\" | cut -d: -f2")
	sysinfo.HardwareUUID = GetCommandOutput("system_profiler SPHardwareDataType | grep \"Hardware UUID\" | cut -d: -f2")

	return sysinfo
}

// CalculateScore returns the compliance score for this system
func CalculateScore(ruleCount int, failCount int) int {
	if ruleCount == 0 { return 0}
	return int(float64(ruleCount-failCount) / float64(ruleCount) * 100.0)
}

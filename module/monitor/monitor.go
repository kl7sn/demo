package monitor

import (
	"runtime"

	"github.com/shirou/gopsutil/v3/mem"

	"github.com/kl7sn/apian/pkg/dto"
)

func New(m dto.Monitor) {

}

// ReadMemStats unit is MB
func ReadMemStats() (uint64, float64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// fmt.Printf("TotalAlloc = %v MiB", m.TotalAlloc/1024/1024)

	v, _ := mem.VirtualMemory()
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// fmt.Printf("Total: %v MiB, Free:%v MiB, UsedPercent:%f%%\n", v.Total/1024/1024, v.Free/1024/1024, v.UsedPercent)

	return m.TotalAlloc / 1024 / 1024, v.UsedPercent
}

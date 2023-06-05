package monitor

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/kl7sn/apian/pkg/dto"
)

func New(m dto.Monitor) {

}

// ReadMemStats unit is MB
func ReadMemStats() (uint64, float64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// fmt.Printf("TotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	usedPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(m.TotalAlloc)/float64(m.Sys)), 64)
	return m.TotalAlloc / 1024 / 1024, usedPercent
}

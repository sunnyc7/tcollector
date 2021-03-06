package collectors

import (
	"github.com/StackExchange/tcollector/opentsdb"
	"github.com/StackExchange/wmi"
)

func init() {
	collectors = append(collectors, &IntervalCollector{F: c_simple_mem_windows})
}

// Memory Needs to be expanded upon. Should be deeper in utilization (what is
// cache, etc.) as well as saturation (i.e., paging activity). Lot of that is in
// Win32_PerfRawData_PerfOS_Memory. Win32_Operating_System's units are KBytes.

func c_simple_mem_windows() opentsdb.MultiDataPoint {
	var dst []Win32_OperatingSystem
	var q = wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		l.Println("simple_mem:", err)
		return nil
	}
	var md opentsdb.MultiDataPoint
	for _, v := range dst {
		Add(&md, "mem.virtual.total", v.TotalVirtualMemorySize*1024, nil)
		Add(&md, "mem.virtual.free", v.FreeVirtualMemory*1024, nil)
		Add(&md, "mem.physical.total", v.TotalVisibleMemorySize*1024, nil)
		Add(&md, "mem.physical.free", v.FreePhysicalMemory*1024, nil)
	}
	return md
}

type Win32_OperatingSystem struct {
	FreePhysicalMemory     uint64
	FreeVirtualMemory      uint64
	TotalVirtualMemorySize uint64
	TotalVisibleMemorySize uint64
}

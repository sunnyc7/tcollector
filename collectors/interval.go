package collectors

import (
	"reflect"
	"runtime"
	"time"

	"github.com/StackExchange/tcollector/opentsdb"
)

type IntervalCollector struct {
	F        func() opentsdb.MultiDataPoint
	Interval time.Duration
}

func (c *IntervalCollector) Run(dpchan chan<- *opentsdb.DataPoint) {
	for {
		interval := c.Interval
		if interval == 0 {
			interval = DEFAULT_FREQ
		}
		next := time.After(interval)
		md := c.F()
		for _, dp := range md {
			dpchan <- dp
		}
		<-next
	}
}

func (c *IntervalCollector) Name() string {
	v := runtime.FuncForPC(reflect.ValueOf(c.F).Pointer())
	return v.Name()
}

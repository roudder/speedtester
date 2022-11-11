package speedtester

import "fmt"

const (
	Fast      string = "fast.com"
	Speedtest string = "speedtest.net"
)

// Measurer common interface for all measurement's functions
type Measurer interface {
	Measure() (*Measurement, error)
}

// Measurement measure's result of network speed
type Measurement struct {
	Resource                   string
	DownloadSpeed, UploadSpeed float64
}

func (measurement Measurement) String() string {
	return fmt.Sprintf("Resource: %v, Download speed: %f Mbps, Upload speed: %f Mbps", measurement.Resource, measurement.DownloadSpeed, measurement.UploadSpeed)
}

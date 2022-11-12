package speedtester

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

type MeasureSpeedtest struct{}

func NewMeasureSpeedtest() *MeasureSpeedtest {
	return &MeasureSpeedtest{}
}

// Measure measures speedtest network speed
func (st *MeasureSpeedtest) Measure() (*Measurement, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, fmt.Errorf("FetchUserInfo fails with: %w", err)
	}
	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return nil, fmt.Errorf("FetchServerList fails with: %w", err)
	}
	if len(serverList.Servers) == 0 {
		return nil, fmt.Errorf("there are no available servers: %w", err)
	}
	// TODO: solution is not ideally, sometimes the closest server isn't in the great state, necessary to add logic in the case when server is not working.
	// get the closest server
	testServer := serverList.Servers[0]

	if err := testServer.DownloadTest(false); err != nil {
		return nil, fmt.Errorf("DownloadTest fails with: %w", err)
	}
	if err := testServer.UploadTest(false); err != nil {
		return nil, fmt.Errorf("UploadTest fails with: %w", err)
	}

	measurement := &Measurement{Resource: Speedtest,
		DownloadSpeed: testServer.DLSpeed,
		UploadSpeed:   testServer.ULSpeed}

	return measurement, nil
}

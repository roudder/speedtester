# Speedtester
Test the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.

Used libs: 
* Inspired and main ideas from the https://github.com/calin014/speedfast (adapted more extended way)
* Delegates to [showwin/speedtest-go](https://github.com/showwin/speedtest-go) for **speedtest.net**



## Dependency

```
go get github.com/roudder/speedtester
```

### API Usage

```go
package main

import (
	"fmt"
	"speedtester"
)

func main() {
	speedtestMeasurement, err := speedtester.NewMeasureSpeedtest().Measure()
	if err != nil {
		//process error as you want
	}
	fmt.Println(speedtestMeasurement)

	fastMeasurement, err := speedtester.NewMeasurementFast().Measure()
	if err != nil {
		//process error as you want
	}
	fmt.Println(fastMeasurement)
}

```
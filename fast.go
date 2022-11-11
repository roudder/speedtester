package speedtester

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gesquive/fast-cli/fast"
	"golang.org/x/sync/errgroup"
)

const (
	workload      = 8
	payloadSizeMB = 25.0 // download payload is by default 25MB, make upload 25MB also
)

type MeasurementFast struct {
	httpClient *http.Client
}

func NewMeasurementFast() *MeasurementFast {
	return &MeasurementFast{
		httpClient: &http.Client{},
	}
}

// Measure naively attempts to measure network speed by using fast.com's api directly
// because fast-cli and go-fast libraries provide only download speed
func (f *MeasurementFast) Measure() (*Measurement, error) {
	urls := fast.GetDlUrls(1)

	if len(urls) == 0 {
		return nil, errors.New("no server urls available")
	}

	url := urls[0]

	downloadSpeed, err := f.measureNetworkSpeed(download, url)
	if err != nil {
		fmt.Errorf("download measureNetworkSpeed fails with: %w", err)
	}
	uploadSpeed, err := f.measureNetworkSpeed(upload, url)
	if err != nil {
		fmt.Errorf("upload measureNetworkSpeed fails with: %w", err)
	}

	measurement := &Measurement{
		Resource:      Fast,
		DownloadSpeed: downloadSpeed,
		UploadSpeed:   uploadSpeed}

	return measurement, nil
}

func (f *MeasurementFast) measureNetworkSpeed(operation func(httpClient *http.Client, url string) error, url string) (float64, error) {
	eg := errgroup.Group{}

	sTime := time.Now()
	for i := 0; i < workload; i++ {
		eg.Go(func() error {
			return operation(f.httpClient, url)
		})
	}
	if err := eg.Wait(); err != nil {
		return 0, err
	}
	fTime := time.Now()

	return payloadSizeMB * 8 * float64(workload) / fTime.Sub(sTime).Seconds(), nil
}

func download(httpClient *http.Client, url string) error {
	if httpClient == nil {
		return errors.New("download: nil httpClient")
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func upload(httpClient *http.Client, uri string) error {
	if httpClient == nil {
		return errors.New("download: nil httpClient")
	}

	v := url.Values{}
	v.Add("content", strings.Repeat("0123456789", payloadSizeMB*1024*1024/10))
	resp, err := httpClient.PostForm(uri, v)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

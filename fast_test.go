package speedtester

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFastMeasure(t *testing.T) {
	fast := NewMeasurementFast()
	result, err := fast.Measure()
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func BenchmarkTestFastMeasure(b *testing.B) {
	fast := NewMeasurementFast()
	for i := 0; i < b.N; i++ {
		fast.Measure()
	}
}

func TestDownload(t *testing.T) {
	t.Run("http get error", func(t *testing.T) {
		httpServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusBadRequest)
		}))
		defer func() { httpServer.Close() }()

		err := download(httpServer.Client(), "testurl")
		assert.Error(t, err)
	})
	t.Run("nil http client", func(t *testing.T) {
		err := download(nil, "testurl")
		assert.Error(t, err)
	})

}

func TestUpload(t *testing.T) {
	t.Run("http get error", func(t *testing.T) {
		httpServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusBadRequest)
		}))
		defer func() { httpServer.Close() }()

		err := upload(httpServer.Client(), "testurl")
		assert.Error(t, err)
	})
	t.Run("nil http client", func(t *testing.T) {
		err := upload(nil, "testurl")
		assert.Error(t, err)
	})
}

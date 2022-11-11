package speedtester

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpeedtestMeasure(t *testing.T) {
	st := NewMeasureSpeedtest()
	result, err := st.Measure()
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func BenchmarkSpeedtestMeasure(b *testing.B) {
	st := NewMeasureSpeedtest()
	for i := 0; i < b.N; i++ {
		st.Measure()
	}
}

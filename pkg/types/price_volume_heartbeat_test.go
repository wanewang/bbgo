package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/wanewang/bbgo/pkg/fixedpoint"
)

func TestPriceHeartBeat_Update(t *testing.T) {
	hb := PriceHeartBeat{}
	updated, err := hb.Update(PriceVolume{Price: fixedpoint.NewFromFloat(22.0), Volume: fixedpoint.NewFromFloat(100.0)}, time.Minute)
	assert.NoError(t, err)
	assert.True(t, updated)

	updated, err = hb.Update(PriceVolume{Price: fixedpoint.NewFromFloat(22.0), Volume: fixedpoint.NewFromFloat(100.0)}, time.Minute)
	assert.NoError(t, err)
	assert.False(t, updated, "should not be updated when pv is not changed")

	updated, err = hb.Update(PriceVolume{Price: fixedpoint.NewFromFloat(23.0), Volume: fixedpoint.NewFromFloat(100.0)}, time.Minute)
	assert.NoError(t, err)
	assert.True(t, updated, "should be updated when the price is changed")

	updated, err = hb.Update(PriceVolume{Price: fixedpoint.NewFromFloat(23.0), Volume: fixedpoint.NewFromFloat(200.0)}, time.Minute)
	assert.NoError(t, err)
	assert.True(t, updated, "should be updated when the volume is changed")
}

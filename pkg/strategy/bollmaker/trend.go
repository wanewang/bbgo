package bollmaker

import "github.com/wanewang/bbgo/pkg/indicator"

type PriceTrend string

const (
	NeutralTrend PriceTrend = "neutral"
	UpTrend      PriceTrend = "upTrend"
	DownTrend    PriceTrend = "downTrend"
	UnknownTrend PriceTrend = "unknown"
)

func detectPriceTrend(inc *indicator.BOLL, price float64) PriceTrend {
	if inBetween(price, inc.LastDownBand(), inc.LastUpBand()) {
		return NeutralTrend
	}

	if price < inc.LastDownBand() {
		return DownTrend
	}

	if price > inc.LastUpBand() {
		return UpTrend
	}

	return UnknownTrend
}

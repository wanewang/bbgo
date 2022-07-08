package backtest

import (
	"github.com/wanewang/bbgo/pkg/bbgo"
	"github.com/wanewang/bbgo/pkg/types"
)

type ExchangeDataSource struct {
	C        chan types.KLine
	Exchange *Exchange
	Session  *bbgo.ExchangeSession
}

// Code generated by "callbackgen -type Stream -interface"; DO NOT EDIT.

package okex

import (
	"github.com/wanewang/bbgo/pkg/exchange/okex/okexapi"
)

func (s *Stream) OnCandleEvent(cb func(candle Candle)) {
	s.candleEventCallbacks = append(s.candleEventCallbacks, cb)
}

func (s *Stream) EmitCandleEvent(candle Candle) {
	for _, cb := range s.candleEventCallbacks {
		cb(candle)
	}
}

func (s *Stream) OnBookEvent(cb func(book BookEvent)) {
	s.bookEventCallbacks = append(s.bookEventCallbacks, cb)
}

func (s *Stream) EmitBookEvent(book BookEvent) {
	for _, cb := range s.bookEventCallbacks {
		cb(book)
	}
}

func (s *Stream) OnEvent(cb func(event WebSocketEvent)) {
	s.eventCallbacks = append(s.eventCallbacks, cb)
}

func (s *Stream) EmitEvent(event WebSocketEvent) {
	for _, cb := range s.eventCallbacks {
		cb(event)
	}
}

func (s *Stream) OnAccountEvent(cb func(account okexapi.Account)) {
	s.accountEventCallbacks = append(s.accountEventCallbacks, cb)
}

func (s *Stream) EmitAccountEvent(account okexapi.Account) {
	for _, cb := range s.accountEventCallbacks {
		cb(account)
	}
}

func (s *Stream) OnOrderDetailsEvent(cb func(orderDetails []okexapi.OrderDetails)) {
	s.orderDetailsEventCallbacks = append(s.orderDetailsEventCallbacks, cb)
}

func (s *Stream) EmitOrderDetailsEvent(orderDetails []okexapi.OrderDetails) {
	for _, cb := range s.orderDetailsEventCallbacks {
		cb(orderDetails)
	}
}

type StreamEventHub interface {
	OnCandleEvent(cb func(candle Candle))

	OnBookEvent(cb func(book BookEvent))

	OnEvent(cb func(event WebSocketEvent))

	OnAccountEvent(cb func(account okexapi.Account))

	OnOrderDetailsEvent(cb func(orderDetails []okexapi.OrderDetails))
}

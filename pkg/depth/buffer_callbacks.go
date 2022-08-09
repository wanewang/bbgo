// Code generated by "callbackgen -type Buffer"; DO NOT EDIT.

package depth

import (
	"github.com/wanewang/bbgo/pkg/types"
)

func (b *Buffer) OnReset(cb func()) {
	b.resetCallbacks = append(b.resetCallbacks, cb)
}

func (b *Buffer) EmitReset() {
	for _, cb := range b.resetCallbacks {
		cb()
	}
}

func (b *Buffer) OnReady(cb func(snapshot types.SliceOrderBook, updates []Update)) {
	b.readyCallbacks = append(b.readyCallbacks, cb)
}

func (b *Buffer) EmitReady(snapshot types.SliceOrderBook, updates []Update) {
	for _, cb := range b.readyCallbacks {
		cb(snapshot, updates)
	}
}

func (b *Buffer) OnPush(cb func(update Update)) {
	b.pushCallbacks = append(b.pushCallbacks, cb)
}

func (b *Buffer) EmitPush(update Update) {
	for _, cb := range b.pushCallbacks {
		cb(update)
	}
}

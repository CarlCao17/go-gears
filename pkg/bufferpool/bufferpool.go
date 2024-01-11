package bufferpool

import (
	"bytes"
	"sync"
)

type Bytes struct {
	p sync.Pool
}

func NewBytesPool() *Bytes {
	return &Bytes{
		p: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (b *Bytes) Get() *bytes.Buffer {
	return b.p.Get().(*bytes.Buffer)
}

func (b *Bytes) Put(x *bytes.Buffer) {
	b.p.Put(x)
}

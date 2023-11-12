package cache

import (
	"sync"
)

type Req struct {
	Params   []interface{}
	Response chan<- Result
}

type Result struct { // output param
	Value interface{}
	Err   error
}

type GenKeyFunc func(args ...interface{}) (key string, err error)
type CachedFunc func(args ...interface{}) (val interface{}, err error)

type Memo struct {
	requests chan Req
	mu       sync.Mutex
	cache    map[string]*entry
}

func New(genKeyFunc GenKeyFunc, callFunc CachedFunc) *Memo {
	memo := &Memo{requests: make(chan Req, 16), cache: make(map[string]*entry)}
	go memo.server(genKeyFunc, callFunc)
	return memo
}

func New2(genKeyFunc GenKeyFunc, callFunc CachedFunc) *Memo {
	memo := &Memo{requests: make(chan Req, 16), cache: make(map[string]*entry)}
	go memo.server2(genKeyFunc, callFunc)
	return memo
}

func (m *Memo) Get(args ...interface{}) (val interface{}, err error) {
	response := make(chan Result)
	m.requests <- Req{
		Params:   args,
		Response: response,
	}
	res := <-response
	return res.Value, res.Err
}

func (m *Memo) Close() {
	close(m.requests)
}

type entry struct {
	res   Result
	ready chan struct{}
}

func (m *Memo) server(genKeyFunc GenKeyFunc, callFunc CachedFunc) {
	for req := range m.requests {
		key, err := genKeyFunc(req.Params...)
		if err != nil {
			go directCall(callFunc, req)
		}
		e := m.cache[key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			m.cache[key] = e
			go e.call(callFunc, req.Params...)
		}
		go e.deliver(req.Response)
	}
	//log.Println("server down")
}

func (m *Memo) server2(keyFunc GenKeyFunc, callFunc CachedFunc) {
	for req := range m.requests {
		go m.handle(keyFunc, callFunc, req)
	}
}

func directCall(callFunc CachedFunc, req Req) {
	var res Result
	res.Value, res.Err = callFunc(req.Params)
	req.Response <- res
	return
}

func (e *entry) deliver(response chan<- Result) {
	<-e.ready
	response <- e.res
}

func (e *entry) call(callFunc CachedFunc, params ...interface{}) {
	e.res.Value, e.res.Err = callFunc(params...)
	close(e.ready)
}

func (m *Memo) handle(genKeyFunc GenKeyFunc, callFunc CachedFunc, req Req) {
	key, err := genKeyFunc(req.Params...)
	if err != nil {
		var res Result
		res.Value, res.Err = callFunc(req.Params)
		req.Response <- res
		return
	}
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()
		go e.call(callFunc, req.Params...)
	} else {
		m.mu.Unlock()
	}
	go e.deliver(req.Response)
}

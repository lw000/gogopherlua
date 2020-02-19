package pipe

import lua "github.com/yuin/gopher-lua"

type Result []lua.LValue

func (r *Result) Add(value lua.LValue) {
	*r = append(*r, value)
}

func (r Result) Len() int {
	return len(r)
}

func (r Result) Get(index int) lua.LValue {
	return r[index]
}

func (r Result) ForEach(fn func(index int, v lua.LValue)) {
	for i, v := range r {
		fn(i, v)
	}
}

type Request struct {
	p    lua.P
	args []lua.LValue
}

func (r *Request) NRet() int {
	return r.p.NRet
}

func (r *Request) P() lua.P {
	return r.p
}

func (r *Request) Args() []lua.LValue {
	return r.args
}

type Data struct {
	Id       int
	request  *Request
	response chan *Result
}

func (p *Data) Request() *Request {
	return p.request
}

func (p *Data) SetResponse(response *Result) {
	if p.response != nil {
		p.response <- response
	}
}

type LuaPipe struct {
	input chan *Data
}

func New() *LuaPipe {
	return &LuaPipe{
		input: make(chan *Data, 1024),
	}
}

func (l *LuaPipe) Call(p lua.P, response chan *Result, args ...lua.LValue) {
	l.input <- &Data{request: &Request{p: p, args: args}, response: response}
}

func (l *LuaPipe) Wait() <-chan *Data {
	return l.input
}

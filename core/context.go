package core

import (
	"context"
	"time"
)

type Context struct {
	*Config
	ctx     context.Context
	ApiAddr string
	Signer  Signer
}

func NewContext(ctx context.Context, cnf *Config, apiAddr string, sign Signer) *Context {
	return &Context{cnf, ctx, apiAddr, sign}
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.ctx.Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.ctx.Done()
}

func (ctx *Context) Err() error {
	return ctx.ctx.Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.ctx.Value(key)
}

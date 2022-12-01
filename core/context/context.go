package context

import (
	"time"

	"github.com/kaiiak/yunzhanghu/core/config"
	"github.com/kaiiak/yunzhanghu/core/credential"
)

type Context struct {
	*config.Config
	ApiAddr string
	Signer  credential.Signer
}

func (ctx Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ctx Context) Done() <-chan struct{} {
	return nil
}

func (ctx Context) Err() error {
	return nil
}

func (ctx Context) Value(key interface{}) interface{} {
	return nil
}

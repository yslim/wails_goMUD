package app

import (
    "context"
)

type MyApp struct {
    ctx context.Context // wails context
}

var _app = &MyApp{}

func App() *MyApp {
    return _app
}

// Ctx get wails context
func (a *MyApp) Ctx() context.Context {
    return a.ctx
}

// SetCtx set wails context
func (a *MyApp) SetCtx(ctx context.Context) *MyApp {
    a.ctx = ctx
    return a
}

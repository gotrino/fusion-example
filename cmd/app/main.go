package main

import (
	"context"
	"github.com/gotrino/fusion-example/turbines"
	_ "github.com/gotrino/fusion-rt-wasmjs"
	"github.com/gotrino/fusion/runtime"
	"github.com/gotrino/fusion/spec/app"
)

type MyApp struct {
}

func (a MyApp) Compose(ctx context.Context) app.Application {
	return app.Application{
		Title: "Mistral",
		Activities: []app.ActivityComposer{
			&turbines.Overview{},
			&turbines.Details{},
		},
		Authentication: app.Bearer{},
	}
}

func main() {
	runtime.MustStart("wasm/js", MyApp{})
}

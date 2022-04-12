package main

import (
	"context"
	"github.com/gotrino/fusion-example/proc"
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
			&proc.Overview{},
			&proc.Details{},
		},
		Connection: app.Connection{
			Scheme: "http",
			Host:   "localhost",
			Port:   8081,
		},
		Authentication: app.HardcodedBearer{
			Token: "3852a83d-4c2d-4e1f-b779-43e17f523703",
		},
	}
}

func main() {
	runtime.MustStart("wasm/js", MyApp{})
}

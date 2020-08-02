package cmd

import (
	"github.com/spiral/endure"
	"github.com/spiral/roadrunner/plugins/env"
	"github.com/spiral/roadrunner/plugins/gzip"
	"github.com/spiral/roadrunner/plugins/headers"
	"github.com/spiral/roadrunner/plugins/health"
	"github.com/spiral/roadrunner/plugins/http"
	"github.com/spiral/roadrunner/plugins/limit"
	"github.com/spiral/roadrunner/plugins/metrics"
	"github.com/spiral/roadrunner/plugins/reload"
	"github.com/spiral/roadrunner/plugins/rpc"
	"github.com/spiral/roadrunner/plugins/static"
)

func main() {
	container, err := endure.NewContainer(endure.DebugLevel)
	if err != nil {
		panic(err)
	}

	err = container.Register(&env.Service{})
	if err != nil {
		panic(err)
	}

	err = container.Register(&rpc.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&http.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&metrics.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&headers.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&static.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&limit.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&health.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&gzip.Service{})
	if err != nil {
		panic(err)
	}
	err = container.Register(&reload.Service{})
	if err != nil {
		panic(err)
	}


	err = container.Init()
	res, err := container.Serve()

	for {
		select {
		case er := <-res:
			println(er)
		}
	}
}

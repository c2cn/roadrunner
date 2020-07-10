package util

import (
	"errors"
	"net/rpc"

	"github.com/spiral/roadrunner/plugins"
	rrpc "github.com/spiral/roadrunner/plugins/rpc"
)

// RPCClient returns RPC client associated with given rr service container.
func RPCClient(container plugins.Container) (*rpc.Client, error) {
	svc, st := container.Get(rrpc.ID)
	if st < plugins.StatusOK {
		return nil, errors.New("RPC service is not configured")
	}

	return svc.(*rrpc.Service).Client()
}

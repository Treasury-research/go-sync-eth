package rpc

import (
	"github.com/chain5j/chain5j-pkg/network/rpc"
)

func NewRpc(host string) (JsonRpc, error) {
	return rpc.Dial(host)
}

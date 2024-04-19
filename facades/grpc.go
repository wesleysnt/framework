package facades

import (
	"github.com/wesleysnt/framework/contracts/grpc"
)

func Grpc() grpc.Grpc {
	return App().MakeGrpc()
}

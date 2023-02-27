package rpc

type RpcServer interface {
	Run(address string)
}

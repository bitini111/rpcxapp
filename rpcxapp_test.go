package rpcxapp

import (
	"fmt"
	"github.com/bitini111/rpcx/server"
	"testing"
)

type RPC struct {
	closed bool
}

func NewRPCApi(closed bool) *RPC {
	return &RPC{
		closed: closed,
	}
}

func Stop(server *server.Server) {

	//关闭接受RPC服务(RPCX收到信号之后没有自动关闭新请求的受理，因此这里自己通过该变量来控制)
	fmt.Println("Server is ShutDown")

}

func TestRun(t *testing.T) {
	fmt.Println("开始测试")
	Run(NewRPCApi(false), Stop)
}

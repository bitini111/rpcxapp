package client

import (
	"github.com/bitini111/rpcx/client"
	"github.com/bitini111/rpcx/protocol"
	"github.com/bitini111/rpcx/share"
	"time"
)

type RouteCallback func(ms map[string]string, args interface{}) []string

//传入的请求参数必须是实现了marshal的接口
//修改说明：在改之前调用XCall时的args参数类型是interface{}，此处添加TheMarshaler之后，就确保了传入的参数必须是实现了Marshal方法的结构体指针。避免了传入结构体的运行时错误。
type TheMarshaler interface {
	Marshal() ([]byte, error)
}

var (
	MyClients          = &Clients{mp: make(map[string]client.XClient), mpcb: make(map[string]RouteCallback)}
	MyClientsWithBytes = &Clients{mp: make(map[string]client.XClient), mpcb: make(map[string]RouteCallback)}
	BasePath           string
	etcAddr            []string
	FailMode           = client.Failfast
	SelectMode         = client.RandomSelect

	//自定义配置，默认编码方式修改为PB，不使用maspack（因为：使用msgpack的时候，pb中修改了字段的话，用到了该结构体的多出服务都需要重启，否则会报如下错误）
	//2020/07/16 17:18:50 server.go:414: WARN : rpcx: failed to handle request: RPC获取userInfo出错！ LoginService Login
	NormalOption = client.Option{
		Retries:        3,
		RPCPath:        share.DefaultRPCPath,
		ConnectTimeout: 10 * time.Second,
		SerializeType:  protocol.ProtoBuffer,
		CompressType:   protocol.None,
		BackupLatency:  10 * time.Millisecond,
	}

	//自定义配置，核心：SerializeType，透传的时候使用
	MyOptions = client.Option{
		Retries:        3,
		RPCPath:        share.DefaultRPCPath,
		ConnectTimeout: 10 * time.Second,
		SerializeType:  protocol.SerializeNone, //此处与默认配置不一样！
		CompressType:   protocol.None,
		BackupLatency:  10 * time.Millisecond,
	}
)
